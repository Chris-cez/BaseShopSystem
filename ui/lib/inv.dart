import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:login_bss/cli.dart';
import 'package:login_bss/crud_template.dart';
import 'package:login_bss/g.dart';
import 'package:login_bss/main.dart';

class I$ extends Source {
  @override
  Future<void> get create async {
    try {
      Response resp = await api.post(
        '/api/sale/draft',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      final invoiceId = resp.data['invoice_id'];
      if (navigatorKey.currentContext != null) {
        await Navigator.of(navigatorKey.currentContext!).push(
          MaterialPageRoute(
            builder: (_) => AddItemsToDraftScreen(invoiceId: invoiceId),
          ),
        );
      }
    } catch (x) {
      throw 'Erro ao criar venda!';
    }
  }

  @override
  Future<void> get delete async {}

  @override
  List<String> get headers => [
        'ID',
        'Número',
        'Cliente',
        'Total',
        'Método Pag.',
        'Desconto',
        'Observação',
        'Chave de Acesso',
      ];

  @override
  List<bool> get show => [
        false,
        false,
        true,
        true,
        true,
        true,
        true,
        false,
      ];

  @override
  Future<void> get fetch async {
    List<List> payload = [];
    try {
      // Busca vendas
      Response resp = await api.get(
        '/api/invoices',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      // Busca clientes
      final clientsResp = await api.get(
        '/api/clients',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      final clients = Map.fromEntries(
        (clientsResp.data['data'] as List)
            .map((c) => MapEntry(c['ID'] ?? c['id'], c['name'] ?? c['Nome'])),
      );
      // Busca métodos de pagamento
      final pmResp = await api.get(
        '/api/payment_methods',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      final paymentMethods = Map.fromEntries(
        (pmResp.data['data'] as List)
            .map((pm) => MapEntry(pm['ID'] ?? pm['id'], pm['name'] ?? pm['Nome'])),
      );
      // Monta a lista de vendas com nomes
      for (Map<String, dynamic> j in resp.data['data']) {
        payload.add([
          j['ID'],
          j['numero'],
          clients[j['client_id']] ?? 'ID ${j['client_id']}',
          j['total'],
          paymentMethods[j['payment_method_id']] ?? 'ID ${j['payment_method_id']}',
          j['discount'],
          j['observation'],
          j['chave_acesso'],
        ]);
      }
    } catch (x) {
      throw 'Erro ao buscar vendas!';
    }
    fetched = [
      ['', '', '', '', '', '', '', ''],
      ...payload,
    ];
  }

  @override
  Future<void> get update async {}

  @override
  List<Fld> get fields => [
        FInt().opt.ned,
        FStr().opt.ned,
        FCli().opt.ned,
        FDbl().opt.ned,
        FInt().opt.ned,
        FDbl().opt.ned,
        FLongStr().opt.ned,
        FStr().opt.ned,
      ];
}

class Ib extends Tb<I$> {
  Ib(super.source);
}

class AddItemsToDraftScreen extends StatefulWidget {
  final String invoiceId;
  const AddItemsToDraftScreen({super.key, required this.invoiceId});

  @override
  State<AddItemsToDraftScreen> createState() => _AddItemsToDraftScreenState();
}

class _AddItemsToDraftScreenState extends State<AddItemsToDraftScreen> {
  Map? selectedProduct;
  Map? selectedClient;
  Map? selectedPaymentMethod;
  final TextEditingController quantityController = TextEditingController();
  List items = [];
  bool loading = false;

  @override
  void initState() {
    super.initState();
    fetchItems();
  }

  Future<List<Map>> fetchProducts() async {
    try {
      final resp = await api.get(
        '/api/products',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      return List<Map>.from(resp.data['data']);
    } catch (_) {
      return [];
    }
  }

  Future<List<Map>> fetchClients() async {
    try {
      final resp = await api.get(
        '/api/clients',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      return List<Map>.from(resp.data['data']);
    } catch (_) {
      return [];
    }
  }

  Future<List<Map>> fetchPaymentMethods() async {
    try {
      final resp = await api.get(
        '/api/payment_methods',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      return List<Map>.from(resp.data['data']);
    } catch (_) {
      return [];
    }
  }

  Future<void> fetchItems() async {
    try {
      final resp = await api.get(
        '/api/sale/items/${widget.invoiceId}',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      // Busca todos os produtos para mapear id -> nome
      final productsResp = await api.get(
        '/api/products',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      final products = Map.fromEntries(
        (productsResp.data['data'] as List)
            .map((p) => MapEntry(p['ID'] ?? p['id'], p['name'] ?? p['Nome'])),
      );
      setState(() {
        items = (resp.data['items'] ?? []).map((item) {
          final prodId = item['product_id'];
          return {
            ...item,
            'product_name': products[prodId] ?? 'ID $prodId',
          };
        }).toList();
      });
    } catch (_) {
      setState(() => items = []);
    }
  }

  Future<void> addItem() async {
    if (selectedProduct == null || quantityController.text.isEmpty) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Selecione um produto e informe a quantidade')),
      );
      return;
    }
    setState(() => loading = true);
    try {
      await api.post(
        '/api/sale/add_item',
        data: {
          'invoice_id': widget.invoiceId,
          'product_id': selectedProduct!['ID'] ?? selectedProduct!['id'],
          'quantity': int.parse(quantityController.text),
        },
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      await fetchItems();
      setState(() {
        selectedProduct = null;
        quantityController.clear();
      });
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erro ao adicionar item')),
      );
    }
    setState(() => loading = false);
  }

  Future<void> finalizeSale() async {
    // Seletores para cliente e método de pagamento
    await showDialog(
      context: context,
      builder: (context) {
        return FutureBuilder(
          future: Future.wait([fetchClients(), fetchPaymentMethods()]),
          builder: (context, AsyncSnapshot<List<dynamic>> snap) {
            if (!snap.hasData) {
              return AlertDialog(content: Center(child: CircularProgressIndicator()));
            }
            final clients = snap.data![0] as List<Map>;
            final paymentMethods = snap.data![1] as List<Map>;
            return StatefulBuilder(
              builder: (context, setStateDialog) => AlertDialog(
                title: Text('Finalizar Venda'),
                content: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    DropdownButton<Map>(
                      isExpanded: true,
                      value: selectedClient,
                      hint: Text('Selecione o Cliente'),
                      items: clients.map((cli) {
                        return DropdownMenuItem<Map>(
                          value: cli,
                          child: Text('${cli['name'] ?? cli['Nome']} (${cli['id'] ?? cli['ID']})'),
                        );
                      }).toList(),
                      onChanged: (v) => setStateDialog(() => selectedClient = v),
                    ),
                    DropdownButton<Map>(
                      isExpanded: true,
                      value: selectedPaymentMethod,
                      hint: Text('Selecione o Método de Pagamento'),
                      items: paymentMethods.map((pm) {
                        return DropdownMenuItem<Map>(
                          value: pm,
                          child: Text('${pm['name'] ?? pm['Nome']} (${pm['id'] ?? pm['ID']})'),
                        );
                      }).toList(),
                      onChanged: (v) => setStateDialog(() => selectedPaymentMethod = v),
                    ),
                  ],
                ),
                actions: [
                  TextButton(
                    onPressed: () => Navigator.of(context).pop(),
                    child: Text('Cancelar'),
                  ),
                  ElevatedButton(
                    onPressed: selectedClient != null && selectedPaymentMethod != null
                        ? () => Navigator.of(context).pop(true)
                        : null,
                    child: Text('Finalizar'),
                  ),
                ],
              ),
            );
          },
        );
      },
    );
    if (selectedClient != null && selectedPaymentMethod != null) {
      setState(() => loading = true);
      try {
        await api.post(
          '/api/sale/finalize',
          data: {
            'invoice_id': widget.invoiceId,
            'client_id': selectedClient!['ID'] ?? selectedClient!['id'],
            'payment_method_id': selectedPaymentMethod!['ID'] ?? selectedPaymentMethod!['id'],
          },
          options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
        );
        if (mounted) {
          Navigator.of(context).pop(true);
        }
      } catch (e) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Erro ao finalizar venda')),
        );
      }
      setState(() => loading = false);
    }
  }

  Future<void> selectProduct() async {
    final products = await fetchProducts();
    final result = await showDialog<Map>(
      context: context,
      builder: (context) => SimpleDialog(
        title: Text('Selecione o Produto'),
        children: products
            .map((prod) => SimpleDialogOption(
                  onPressed: () => Navigator.pop(context, prod),
                  child: Text('${prod['name'] ?? prod['Nome']} (${prod['id'] ?? prod['ID']})'),
                ))
            .toList(),
      ),
    );
    if (result != null) setState(() => selectedProduct = result);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Itens da Venda')),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          children: [
            Row(
              children: [
                Expanded(
                  child: InkWell(
                    onTap: loading ? null : selectProduct,
                    child: InputDecorator(
                      decoration: InputDecoration(
                        labelText: 'Produto',
                        border: OutlineInputBorder(),
                      ),
                      child: Text(
                        selectedProduct == null
                            ? 'Selecione'
                            : '${selectedProduct!['name'] ?? selectedProduct!['Nome']} (${selectedProduct!['id'] ?? selectedProduct!['ID']})',
                      ),
                    ),
                  ),
                ),
                SizedBox(width: 8),
                Expanded(
                  child: TextField(
                    controller: quantityController,
                    decoration: InputDecoration(labelText: 'Quantidade'),
                    keyboardType: TextInputType.number,
                  ),
                ),
                SizedBox(width: 8),
                ElevatedButton(
                  onPressed: loading ? null : addItem,
                  child: loading
                      ? SizedBox(width: 16, height: 16, child: CircularProgressIndicator(strokeWidth: 2))
                      : Text('Adicionar'),
                ),
              ],
            ),
            SizedBox(height: 16),
            Expanded(
              child: items.isEmpty
                  ? Center(child: Text('Nenhum item adicionado'))
                  : ListView.builder(
                      itemCount: items.length,
                      itemBuilder: (context, i) {
                        final item = items[i];
                        return ListTile(
                          title: Text('Produto: ${item['product_name']}'),
                          subtitle: Text('Quantidade: ${item['quantity']}'),
                        );
                      },
                    ),
            ),
            ElevatedButton(
              onPressed: loading ? null : finalizeSale,
              child: Text('Finalizar Venda'),
            ),
          ],
        ),
      ),
    );
  }
  }