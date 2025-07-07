import 'package:flutter/material.dart';
import 'package:dio/dio.dart';
import 'g.dart';

class VendasPage extends StatefulWidget {
  const VendasPage();

  @override
  State<VendasPage> createState() => _VendasPageState();
}

class _VendasPageState extends State<VendasPage> {
  List<Venda> vendas = [];
  bool loading = false;
  String? erro;

  @override
  void initState() {
    super.initState();
    _carregarVendas();
  }

  Future<void> _carregarVendas() async {
    setState(() { loading = true; erro = null; });
    try {
      vendas = await VendaService(api, key.value ?? '').listarVendas();
    } catch (e) {
      erro = e.toString();
    }
    setState(() { loading = false; });
  }

  Future<void> _criarVenda() async {
    setState(() { loading = true; });
    try {
      await VendaService(api, key.value ?? '').criarVenda();
      await _carregarVendas();
    } catch (e) {
      setState(() { erro = e.toString(); });
    }
    setState(() { loading = false; });
  }

  void _abrirDetalhe(Venda venda) async {
    final atualizado = await showModalBottomSheet<bool>(
      context: context,
      isScrollControlled: true,
      builder: (_) => VendaDetalhe(venda: venda, onAtualizar: _carregarVendas),
    );
    if (atualizado == true) {
      await _carregarVendas();
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Vendas'),
        actions: [
          IconButton(
            icon: const Icon(Icons.refresh),
            onPressed: _carregarVendas,
          ),
          IconButton(
            icon: const Icon(Icons.add),
            onPressed: _criarVenda,
            tooltip: 'Nova venda',
          ),
        ],
      ),
      body: loading
          ? const Center(child: CircularProgressIndicator())
          : erro != null
              ? Center(child: Text('Erro: $erro'))
              : vendas.isEmpty
                  ? const Center(child: Text('Nenhuma venda encontrada.'))
                  : ListView.builder(
                      itemCount: vendas.length,
                      itemBuilder: (context, idx) {
                        final v = vendas[idx];
                        return Card(
                          margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
                          child: ListTile(
                            title: Text('Cliente: ${v.cliente}'),
                            subtitle: Text('Total: R\$ ${v.total.toStringAsFixed(2)}\nMétodo: ${v.metodoPagamento}\nStatus: ${v.finalizada ? 'Finalizada' : 'Em aberto'}'),
                            trailing: const Icon(Icons.chevron_right),
                            onTap: () => _abrirDetalhe(v),
                          ),
                        );
                      },
                    ),
    );
  }
}

class VendaDetalhe extends StatefulWidget {
  final Venda venda;
  final VoidCallback onAtualizar;
  const VendaDetalhe({required this.venda, required this.onAtualizar});

  @override
  State<VendaDetalhe> createState() => _VendaDetalheState();
}

class _VendaDetalheState extends State<VendaDetalhe> {
  bool loading = false;
  String? erro;
  List<ItemVenda> itens = [];
  List<DropdownMenuItem<String>> produtos = [];
  String? produtoSelecionado;
  int quantidade = 1;
  List<DropdownMenuItem<String>> clientes = [];
  String? clienteSelecionado;
  List<DropdownMenuItem<String>> metodos = [];
  String? metodoSelecionado;

  @override
  void initState() {
    super.initState();
    itens = widget.venda.itens;
    clienteSelecionado = _getClienteIdByNome(widget.venda.cliente);
    metodoSelecionado = _getMetodoIdByNome(widget.venda.metodoPagamento);
    _carregarDropdowns();
  }

  String? _getClienteIdByNome(String nome) {
    // Será preenchido corretamente após _carregarDropdowns
    return null;
  }
  String? _getMetodoIdByNome(String nome) {
    // Será preenchido corretamente após _carregarDropdowns
    return null;
  }

  Future<void> _carregarDropdowns() async {
    final service = VendaService(api, key.value ?? '');
    final prods = await service._produtos();
    final clis = await service._clientes();
    final mets = await service._metodosPagamento();
    setState(() {
      produtos = prods.entries.map((e) => DropdownMenuItem(value: e.key, child: Text(e.value))).toList();
      clientes = clis.entries.map((e) => DropdownMenuItem(value: e.key, child: Text(e.value))).toList();
      metodos = mets.entries.map((e) => DropdownMenuItem(value: e.key, child: Text(e.value))).toList();
      // Seleciona o cliente/metodo correto se já existe na venda
      clienteSelecionado = clis.entries.firstWhere(
        (e) => e.value == widget.venda.cliente,
        orElse: () => MapEntry('', ''),
      ).key.isNotEmpty ? clis.entries.firstWhere((e) => e.value == widget.venda.cliente).key : null;
      metodoSelecionado = mets.entries.firstWhere(
        (e) => e.value == widget.venda.metodoPagamento,
        orElse: () => MapEntry('', ''),
      ).key.isNotEmpty ? mets.entries.firstWhere((e) => e.value == widget.venda.metodoPagamento).key : null;
    });
  }

  Future<void> _adicionarItem() async {
    if (produtoSelecionado == null || quantidade < 1) return;
    setState(() { loading = true; erro = null; });
    try {
      await VendaService(api, key.value ?? '').adicionarItem(widget.venda.numero, produtoSelecionado!, quantidade);
      final prods = await VendaService(api, key.value ?? '')._produtos();
      itens = await VendaService(api, key.value ?? '').listarItens(widget.venda.numero, prods);
      setState(() {});
      widget.onAtualizar();
    } catch (e) {
      setState(() { erro = e.toString(); });
    }
    setState(() { loading = false; });
  }

  Future<void> _finalizarVenda() async {
    if (clienteSelecionado == null || metodoSelecionado == null) return;
    setState(() { loading = true; erro = null; });
    try {
      await VendaService(api, key.value ?? '').finalizarVenda(widget.venda.numero, clienteSelecionado!, metodoSelecionado!);
      Navigator.pop(context, true); // Sinaliza que precisa atualizar a lista
    } catch (e) {
      setState(() { erro = e.toString(); });
    }
    setState(() { loading = false; });
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.only(
        left: 16,
        right: 16,
        top: 16,
        bottom: MediaQuery.of(context).viewInsets.bottom + 16,
      ),
      child: SingleChildScrollView(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text('Cliente: ${widget.venda.cliente}', style: const TextStyle(fontWeight: FontWeight.bold)),
            Text('Método: ${widget.venda.metodoPagamento}'),
            Text('Status: ${widget.venda.finalizada ? 'Finalizada' : 'Em aberto'}'),
            const SizedBox(height: 16),
            const Text('Itens da venda:', style: TextStyle(fontWeight: FontWeight.bold)),
            ...itens.map((item) => ListTile(
                  title: Text(item.produto),
                  subtitle: Text('Qtd: ${item.quantidade}  |  Valor: R\$ ${item.valorUnitario.toStringAsFixed(2)}'),
                )),
            if (!widget.venda.finalizada) ...[
              const Divider(),
              Row(
                children: [
                  Expanded(
                    child: DropdownButtonFormField<String>(
                      value: produtoSelecionado,
                      items: produtos,
                      onChanged: (v) => setState(() => produtoSelecionado = v),
                      decoration: const InputDecoration(labelText: 'Produto'),
                    ),
                  ),
                  const SizedBox(width: 8),
                  SizedBox(
                    width: 80,
                    child: TextFormField(
                      initialValue: '1',
                      keyboardType: TextInputType.number,
                      decoration: const InputDecoration(labelText: 'Qtd'),
                      onChanged: (v) => quantidade = int.tryParse(v) ?? 1,
                    ),
                  ),
                  IconButton(
                    icon: const Icon(Icons.add),
                    onPressed: loading ? null : _adicionarItem,
                  ),
                ],
              ),
              const SizedBox(height: 16),
              DropdownButtonFormField<String>(
                value: clienteSelecionado,
                items: clientes,
                onChanged: (v) => setState(() => clienteSelecionado = v),
                decoration: const InputDecoration(labelText: 'Cliente'),
              ),
              const SizedBox(height: 8),
              DropdownButtonFormField<String>(
                value: metodoSelecionado,
                items: metodos,
                onChanged: (v) => setState(() => metodoSelecionado = v),
                decoration: const InputDecoration(labelText: 'Método de Pagamento'),
              ),
              const SizedBox(height: 16),
              ElevatedButton(
                onPressed: loading ? null : _finalizarVenda,
                child: const Text('Finalizar venda'),
              ),
            ],
            if (erro != null) ...[
              const SizedBox(height: 8),
              Text('Erro: $erro', style: const TextStyle(color: Colors.red)),
            ],
          ],
        ),
      ),
    );
  }
}

// MODELOS E SERVIÇO
class Venda {
  final String numero;
  final String cliente;
  final double total;
  final String metodoPagamento;
  final double desconto;
  final String observacao;
  final bool finalizada;
  final List<ItemVenda> itens;

  Venda({
    required this.numero,
    required this.cliente,
    required this.total,
    required this.metodoPagamento,
    required this.desconto,
    required this.observacao,
    required this.finalizada,
    required this.itens,
  });
}

class ItemVenda {
  final String produto;
  final int quantidade;
  final double valorUnitario;
  final double valorTotal;

  ItemVenda({
    required this.produto,
    required this.quantidade,
    required this.valorUnitario,
    required this.valorTotal,
  });

}

class VendaService {
  final Dio api;
  final String token;
  VendaService(this.api, this.token);

  Future<List<Venda>> listarVendas() async {
    final clientes = await _clientes();
    final metodos = await _metodosPagamento();
    final produtos = await _produtos();
    final resp = await api.get(
      '/api/invoices',
      options: Options(headers: {'Authorization': 'Bearer $token'}),
    );
    List<Venda> vendas = [];
    for (final j in resp.data['data']) {
      final itens = await listarItens(j['numero'], produtos);
      final accessKey = j['access_key']?.toString() ?? '';
      // Usa sempre o campo total_value do backend para o total da venda
      double total = (j['total_value'] ?? 0).toDouble();
      vendas.add(Venda(
        numero: j['numero'] ?? '',
        cliente: clientes[j['client_id']?.toString() ?? ''] ?? '-',
        total: total,
        metodoPagamento: metodos[j['payment_method_id']?.toString() ?? ''] ?? '-',
        desconto: (j['discount'] ?? 0).toDouble(),
        observacao: j['observation'] ?? '',
        finalizada: accessKey.isNotEmpty,
        itens: itens,
      ));
    }
    return vendas;
  }

  Future<List<ItemVenda>> listarItens(String numero, Map<String, String> produtos) async {
    final resp = await api.get(
      '/api/sale/items/$numero',
      options: Options(headers: {'Authorization': 'Bearer $token'}),
    );
    return [
      for (final j in resp.data['items'])
        ItemVenda(
          produto: produtos[j['product_id'].toString()] ?? 'Produto',
          quantidade: (j['quantity'] is int) ? j['quantity'] : int.tryParse(j['quantity'].toString()) ?? 0,
          valorUnitario: (j['price'] is double) ? j['price'] : double.tryParse(j['price'].toString()) ?? 0.0,
          valorTotal: (j['valor_total'] is double) ? j['valor_total'] : double.tryParse(j['valor_total'].toString()) ?? 0.0,
        )
    ];
  }

  Future<String> criarVenda() async {
    final resp = await api.post(
      '/api/sale/draft',
      options: Options(headers: {'Authorization': 'Bearer $token'}),
    );
    return resp.data['invoice_id'] ?? '';
  }

  Future<void> adicionarItem(String numero, String produtoId, int quantidade) async {
    await api.post(
      '/api/sale/add_item',
      data: {
        'invoice_id': numero,
        'product_id': int.parse(produtoId),
        'quantity': quantidade,
      },
      options: Options(headers: {'Authorization': 'Bearer $token'}),
    );
  }

  Future<void> finalizarVenda(String numero, String clienteId, String metodoPagamentoId) async {
    await api.post(
      '/api/sale/finalize',
      data: {
        'invoice_id': numero,
        'client_id': int.parse(clienteId),
        'payment_method_id': int.parse(metodoPagamentoId),
      },
      options: Options(headers: {'Authorization': 'Bearer $token'}),
    );
  }

  // Mocks para nomes (substitua por chamadas reais)
  Future<Map<String, String>> _clientes() async {
    final resp = await api.get(
      '/api/clients',
      options: Options(headers: {'Authorization': 'Bearer $token'}),
    );
    final data = resp.data['data'] as List<dynamic>?;
    if (data == null) return {};
    return {
      for (final c in data)
        (c['ID']?.toString() ?? c['id']?.toString() ?? ''): c['name'] ?? ''
    };
  }

  Future<Map<String, String>> _metodosPagamento() async {
    final resp = await api.get(
      '/api/payment_methods',
      options: Options(headers: {'Authorization': 'Bearer $token'}),
    );
    final data = resp.data['data'] as List<dynamic>?;
    if (data == null) return {};
    return {
      for (final m in data)
        (m['ID']?.toString() ?? m['id']?.toString() ?? ''): m['name'] ?? ''
    };
  }

  // Busca produtos reais do backend
  Future<Map<String, String>> _produtos() async {
    final resp = await api.get(
      '/api/products',
      options: Options(headers: {'Authorization': 'Bearer $token'}),
    );
    final data = resp.data['data'] as List<dynamic>?;
    if (data == null) return {};
    return {
      for (final p in data)
        (p['ID']?.toString() ?? p['id']?.toString() ?? ''): p['name'] ?? ''
    };
  }

}