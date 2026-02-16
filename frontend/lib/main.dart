import 'package:flutter/material.dart';
import 'package:grpc/grpc_web.dart';
import 'proto/poker.pbgrpc.dart';

void main() => runApp(const PokerApp());

class PokerApp extends StatelessWidget {
  const PokerApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Texas Hold\'em Calculator',
      theme: ThemeData(
        colorSchemeSeed: Colors.green,
        useMaterial3: true,
      ),
      home: const PokerHomePage(),
    );
  }
}

class PokerHomePage extends StatefulWidget {
  const PokerHomePage({super.key});

  @override
  State<PokerHomePage> createState() => _PokerHomePageState();
}

class _PokerHomePageState extends State<PokerHomePage> {
  late final _channel = GrpcWebClientChannel.xhr(
    Uri.parse('http://localhost:8080'),
  );
  late final _client = PokerServiceClient(_channel);

  @override
  void dispose() {
    _channel.shutdown();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 3,
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Texas Hold\'em Calculator'),
          bottom: const TabBar(
            tabs: [
              Tab(icon: Icon(Icons.style), text: 'Evaluate'),
              Tab(icon: Icon(Icons.compare_arrows), text: 'Compare'),
              Tab(icon: Icon(Icons.calculate), text: 'Probability'),
            ],
          ),
        ),
        body: TabBarView(
          children: [
            EvaluateTab(client: _client),
            CompareTab(client: _client),
            ProbabilityTab(client: _client),
          ],
        ),
      ),
    );
  }
}

// ─── Card input helper ───

class CardInputField extends StatelessWidget {
  final TextEditingController controller;
  final String label;
  final String hint;

  const CardInputField({
    super.key,
    required this.controller,
    required this.label,
    required this.hint,
  });

  @override
  Widget build(BuildContext context) {
    return TextField(
      controller: controller,
      decoration: InputDecoration(
        labelText: label,
        hintText: hint,
        border: const OutlineInputBorder(),
        helperText: 'Format: Suit(H/D/C/S) + Rank(2-9/T/J/Q/K/A)',
        helperMaxLines: 2,
      ),
      textCapitalization: TextCapitalization.characters,
    );
  }
}

// ─── Evaluate Tab ───

class EvaluateTab extends StatefulWidget {
  final PokerServiceClient client;
  const EvaluateTab({super.key, required this.client});

  @override
  State<EvaluateTab> createState() => _EvaluateTabState();
}

class _EvaluateTabState extends State<EvaluateTab> {
  final _holeCtrl = TextEditingController();
  final _communityCtrl = TextEditingController();
  String _result = '';
  bool _loading = false;

  Future<void> _evaluate() async {
    final hole = _holeCtrl.text.trim().split(RegExp(r'\s+'));
    final community = _communityCtrl.text.trim().split(RegExp(r'\s+'));
    final allCards = [...hole, ...community];

    if (allCards.length != 7) {
      setState(() =>
          _result = 'Need exactly 7 cards (2 hole + 5 community). '
              'Got ${allCards.length}.');
      return;
    }

    setState(() => _loading = true);
    try {
      final resp = await widget.client.evaluateHand(
        EvaluateHandRequest()..cards.addAll(allCards),
      );
      setState(() {
        _result = 'Hand: ${resp.handRank}\n'
            'Best 5: ${resp.bestFive.join(' ')}\n'
            'Rank Value: ${resp.rankValue}';
      });
    } catch (e) {
      setState(() => _result = 'Error: $e');
    } finally {
      setState(() => _loading = false);
    }
  }

  @override
  void dispose() {
    _holeCtrl.dispose();
    _communityCtrl.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      padding: const EdgeInsets.all(24),
      child: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 500),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              Text('Hand Evaluation',
                  style: Theme.of(context).textTheme.headlineSmall),
              const SizedBox(height: 8),
              const Text('Enter 2 hole cards and 5 community cards '
                  'to find the best hand.'),
              const SizedBox(height: 16),
              CardInputField(
                  controller: _holeCtrl,
                  label: 'Hole Cards (2)',
                  hint: 'HA SK'),
              const SizedBox(height: 12),
              CardInputField(
                  controller: _communityCtrl,
                  label: 'Community Cards (5)',
                  hint: 'D2 CT H5 S9 DJ'),
              const SizedBox(height: 16),
              FilledButton(
                onPressed: _loading ? null : _evaluate,
                child: _loading
                    ? const SizedBox(
                        width: 20,
                        height: 20,
                        child: CircularProgressIndicator(strokeWidth: 2))
                    : const Text('Evaluate Hand'),
              ),
              const SizedBox(height: 24),
              if (_result.isNotEmpty)
                Card(
                  child: Padding(
                    padding: const EdgeInsets.all(16),
                    child: Text(_result,
                        style: Theme.of(context).textTheme.bodyLarge),
                  ),
                ),
            ],
          ),
        ),
      ),
    );
  }
}

// ─── Compare Tab ───

class CompareTab extends StatefulWidget {
  final PokerServiceClient client;
  const CompareTab({super.key, required this.client});

  @override
  State<CompareTab> createState() => _CompareTabState();
}

class _CompareTabState extends State<CompareTab> {
  final _p1HoleCtrl = TextEditingController();
  final _p2HoleCtrl = TextEditingController();
  final _communityCtrl = TextEditingController();
  String _result = '';
  bool _loading = false;

  Future<void> _compare() async {
    final p1Hole = _p1HoleCtrl.text.trim().split(RegExp(r'\s+'));
    final p2Hole = _p2HoleCtrl.text.trim().split(RegExp(r'\s+'));
    final community = _communityCtrl.text.trim().split(RegExp(r'\s+'));

    if (p1Hole.length != 2 || p2Hole.length != 2 || community.length != 5) {
      setState(() =>
          _result = 'Need 2 hole cards per player and 5 community cards.');
      return;
    }

    setState(() => _loading = true);
    try {
      final resp = await widget.client.compareHands(
        CompareHandsRequest()
          ..player1Cards.addAll([...p1Hole, ...community])
          ..player2Cards.addAll([...p2Hole, ...community]),
      );
      setState(() {
        _result = '${resp.description}\n\n'
            'Player 1: ${resp.player1Hand.handRank} '
            '(${resp.player1Hand.bestFive.join(' ')})\n'
            'Player 2: ${resp.player2Hand.handRank} '
            '(${resp.player2Hand.bestFive.join(' ')})';
      });
    } catch (e) {
      setState(() => _result = 'Error: $e');
    } finally {
      setState(() => _loading = false);
    }
  }

  @override
  void dispose() {
    _p1HoleCtrl.dispose();
    _p2HoleCtrl.dispose();
    _communityCtrl.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      padding: const EdgeInsets.all(24),
      child: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 500),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              Text('Hand Comparison',
                  style: Theme.of(context).textTheme.headlineSmall),
              const SizedBox(height: 8),
              const Text('Compare two players\' hands with shared '
                  'community cards.'),
              const SizedBox(height: 16),
              CardInputField(
                  controller: _p1HoleCtrl,
                  label: 'Player 1 Hole Cards (2)',
                  hint: 'HA HK'),
              const SizedBox(height: 12),
              CardInputField(
                  controller: _p2HoleCtrl,
                  label: 'Player 2 Hole Cards (2)',
                  hint: 'SA SK'),
              const SizedBox(height: 12),
              CardInputField(
                  controller: _communityCtrl,
                  label: 'Community Cards (5)',
                  hint: 'D2 CT H5 S9 DJ'),
              const SizedBox(height: 16),
              FilledButton(
                onPressed: _loading ? null : _compare,
                child: _loading
                    ? const SizedBox(
                        width: 20,
                        height: 20,
                        child: CircularProgressIndicator(strokeWidth: 2))
                    : const Text('Compare Hands'),
              ),
              const SizedBox(height: 24),
              if (_result.isNotEmpty)
                Card(
                  child: Padding(
                    padding: const EdgeInsets.all(16),
                    child: Text(_result,
                        style: Theme.of(context).textTheme.bodyLarge),
                  ),
                ),
            ],
          ),
        ),
      ),
    );
  }
}

// ─── Probability Tab ───

class ProbabilityTab extends StatefulWidget {
  final PokerServiceClient client;
  const ProbabilityTab({super.key, required this.client});

  @override
  State<ProbabilityTab> createState() => _ProbabilityTabState();
}

class _ProbabilityTabState extends State<ProbabilityTab> {
  final _holeCtrl = TextEditingController();
  final _communityCtrl = TextEditingController();
  int _numOpponents = 1;
  int _iterations = 10000;
  String _result = '';
  bool _loading = false;

  Future<void> _calculate() async {
    final hole = _holeCtrl.text.trim().split(RegExp(r'\s+'));
    if (hole.length != 2) {
      setState(() => _result = 'Need exactly 2 hole cards.');
      return;
    }

    final communityText = _communityCtrl.text.trim();
    final community =
        communityText.isEmpty ? <String>[] : communityText.split(RegExp(r'\s+'));

    if (community.length > 5) {
      setState(() => _result = 'At most 5 community cards.');
      return;
    }

    setState(() => _loading = true);
    try {
      final resp = await widget.client.calculateWinProbability(
        WinProbabilityRequest()
          ..holeCards.addAll(hole)
          ..community.addAll(community)
          ..numOpponents = _numOpponents
          ..iterations = _iterations,
      );
      setState(() {
        _result = 'Win:  ${resp.winProbability.toStringAsFixed(1)}%\n'
            'Tie:  ${resp.tieProbability.toStringAsFixed(1)}%\n'
            'Loss: ${resp.lossProbability.toStringAsFixed(1)}%\n'
            'Simulations: ${resp.iterationsRun}';
      });
    } catch (e) {
      setState(() => _result = 'Error: $e');
    } finally {
      setState(() => _loading = false);
    }
  }

  @override
  void dispose() {
    _holeCtrl.dispose();
    _communityCtrl.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      padding: const EdgeInsets.all(24),
      child: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 500),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              Text('Win Probability',
                  style: Theme.of(context).textTheme.headlineSmall),
              const SizedBox(height: 8),
              const Text('Monte Carlo simulation to estimate your '
                  'winning chances.'),
              const SizedBox(height: 16),
              CardInputField(
                  controller: _holeCtrl,
                  label: 'Your Hole Cards (2)',
                  hint: 'HA DA'),
              const SizedBox(height: 12),
              CardInputField(
                  controller: _communityCtrl,
                  label: 'Community Cards (0-5, optional)',
                  hint: 'D2 CT H5'),
              const SizedBox(height: 16),
              Row(
                children: [
                  Expanded(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text('Opponents: $_numOpponents'),
                        Slider(
                          value: _numOpponents.toDouble(),
                          min: 1,
                          max: 9,
                          divisions: 8,
                          label: '$_numOpponents',
                          onChanged: (v) =>
                              setState(() => _numOpponents = v.round()),
                        ),
                      ],
                    ),
                  ),
                  const SizedBox(width: 16),
                  Expanded(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text('Simulations: $_iterations'),
                        Slider(
                          value: _iterations.toDouble(),
                          min: 1000,
                          max: 100000,
                          divisions: 99,
                          label: '$_iterations',
                          onChanged: (v) =>
                              setState(() => _iterations = v.round()),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 16),
              FilledButton(
                onPressed: _loading ? null : _calculate,
                child: _loading
                    ? const SizedBox(
                        width: 20,
                        height: 20,
                        child: CircularProgressIndicator(strokeWidth: 2))
                    : const Text('Calculate Probability'),
              ),
              const SizedBox(height: 24),
              if (_result.isNotEmpty)
                Card(
                  child: Padding(
                    padding: const EdgeInsets.all(16),
                    child: Text(_result,
                        style: Theme.of(context).textTheme.bodyLarge),
                  ),
                ),
            ],
          ),
        ),
      ),
    );
  }
}
