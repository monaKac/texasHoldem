#!/usr/bin/env bash
set -euo pipefail

# ─── Configuration ────────────────────────────────────────────────────────────
PROJECT_ID="YOUR_PROJECT_ID"        # ← replace with your GCP project ID
REGION="europe-west1"
CLUSTER_NAME="texasholdem-cluster"
REPO_NAME="texasholdem"
TAG="latest"

# Derived
REGISTRY="${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO_NAME}"
BACKEND_IMAGE="${REGISTRY}/backend:${TAG}"
FRONTEND_IMAGE="${REGISTRY}/frontend:${TAG}"

# ─── Preflight checks ────────────────────────────────────────────────────────
if [[ "$PROJECT_ID" == "YOUR_PROJECT_ID" ]]; then
  echo "ERROR: Set PROJECT_ID in this script before running." >&2
  exit 1
fi

for cmd in gcloud docker kubectl; do
  if ! command -v "$cmd" &>/dev/null; then
    echo "ERROR: $cmd is required but not installed." >&2
    exit 1
  fi
done

gcloud config set project "$PROJECT_ID"

# ─── 1. Artifact Registry ────────────────────────────────────────────────────
echo "▸ Ensuring Artifact Registry repo exists…"
if ! gcloud artifacts repositories describe "$REPO_NAME" \
      --location="$REGION" &>/dev/null; then
  gcloud artifacts repositories create "$REPO_NAME" \
    --repository-format=docker \
    --location="$REGION" \
    --description="texasHoldem container images"
fi

# Configure Docker to authenticate with Artifact Registry
gcloud auth configure-docker "${REGION}-docker.pkg.dev" --quiet

# ─── 2. GKE Autopilot cluster ────────────────────────────────────────────────
echo "▸ Ensuring GKE Autopilot cluster exists…"
if ! gcloud container clusters describe "$CLUSTER_NAME" \
      --region="$REGION" &>/dev/null; then
  gcloud container clusters create-auto "$CLUSTER_NAME" \
    --region="$REGION"
fi

# Point kubectl at the cluster
gcloud container clusters get-credentials "$CLUSTER_NAME" --region="$REGION"

# ─── 3. Build & push images ──────────────────────────────────────────────────
echo "▸ Building and pushing backend image…"
docker build --platform linux/amd64 -t "$BACKEND_IMAGE" ./backend
docker push "$BACKEND_IMAGE"

echo "▸ Building and pushing frontend image…"
docker build --platform linux/amd64 -t "$FRONTEND_IMAGE" ./frontend
docker push "$FRONTEND_IMAGE"

# ─── 4. Deploy to GKE ────────────────────────────────────────────────────────
echo "▸ Applying Kubernetes manifests…"

# Substitute image placeholders and apply
sed "s|IMAGE_BACKEND|${BACKEND_IMAGE}|g" k8s/backend.yaml  | kubectl apply -f -
sed "s|IMAGE_FRONTEND|${FRONTEND_IMAGE}|g" k8s/frontend.yaml | kubectl apply -f -
kubectl apply -f k8s/envoy.yaml

# ─── 5. Wait for external IPs ────────────────────────────────────────────────
echo "▸ Waiting for external IPs (this may take a minute)…"

wait_for_ip() {
  local svc="$1"
  local ip=""
  while [[ -z "$ip" || "$ip" == "<pending>" ]]; do
    ip=$(kubectl get svc "$svc" -o jsonpath='{.status.loadBalancer.ingress[0].ip}' 2>/dev/null || true)
    sleep 5
  done
  echo "$ip"
}

FRONTEND_IP=$(wait_for_ip frontend)
ENVOY_IP=$(wait_for_ip envoy)

echo ""
echo "════════════════════════════════════════════"
echo "  Deployment complete!"
echo ""
echo "  Frontend:  http://${FRONTEND_IP}"
echo "  Envoy:     http://${ENVOY_IP}:8080"
echo "════════════════════════════════════════════"
