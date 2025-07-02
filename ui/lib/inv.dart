import 'package:dio/dio.dart';
import 'package:login_bss/cli.dart';
import 'package:login_bss/crud_template.dart';
import 'package:login_bss/g.dart';

class I$ extends Source {

  @override
  Future<void> get create async {
    print(fetched[0]);
    try {
      Response resp = await api.post(
        '/api/sale/draft',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
    } catch (x) {
      throw 'Erro!';
    }
  }

  @override
  Future<void> get delete async {
    // Faz nada
  }

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
      Response resp = await api.get(
        '/api/invoices',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      print(resp.data);
      for (Map<String, dynamic> j in resp.data['data']) {
        payload.add([
          j['ID'],
          j['numero'],
          j['client_id'],
          j['total_value'],
          j['payment_method_id'],
          j['discount'],
          j['observation'],
          j['access_key'],
        ]);
      }
    } catch (x) {
      throw 'Erro!';
    }
    fetched = [
      ['', '', '', '', '', '', '', ''],
      ...payload,
    ];
    print(fetched);
  }

  @override
  Future<void> get update async {
    try {
      Response resp = await api.put(
        '/api/invoices/${temp[0]}',
        data: {
          'numero': temp[1],
          'client_id': temp[2],
          'total_value': temp[3],
          'payment_method_id': temp[4],
          'discount': temp[5],
          'observation': temp[6],
          'access_key': temp[7],
        },
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
    } on DioException catch (a, x) {
      throw 'Erro! ${a.response!.data}';
    }
  }

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