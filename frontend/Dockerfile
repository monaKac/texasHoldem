# ---------- build stage ----------
FROM debian:bookworm-slim AS builder

RUN apt-get update && apt-get install -y --no-install-recommends \
    curl git unzip xz-utils ca-certificates \
    && rm -rf /var/lib/apt/lists/*

ENV FLUTTER_HOME=/opt/flutter
RUN git clone --depth 1 --branch stable https://github.com/flutter/flutter.git $FLUTTER_HOME
ENV PATH="$FLUTTER_HOME/bin:$PATH"
RUN flutter precache --web && flutter doctor -v

WORKDIR /app
COPY . .
RUN flutter pub get && flutter build web --release

# ---------- runtime stage ----------
FROM nginx:1.27-alpine
COPY --from=builder /app/build/web /usr/share/nginx/html
EXPOSE 80
