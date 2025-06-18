import 'package:dio/dio.dart';
import 'package:login_bss/crud_template.dart';
import 'package:login_bss/g.dart';

class P$ extends Source {

  @override
  Future<void> get create async {
    print(fetched[0]);
    try {
      Response resp = await api.post(
        '/api/products',
        data: {
          'code': fetched[0][0],
          'price': fetched[0][1],
          'name': fetched[0][2],
          'gtin': fetched[0][3], // temp[3],
          'um': fetched[0][4],
          'description': fetched[0][5], // temp[5],
          'class_id': fetched[0][6], //fetched[0][6],
          'stock': fetched[0][7], //temp[7],
          'vtribute': fetched[0][8] // temp[8],
        },
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
    } catch (x) {
      throw 'Erro!';
    }
  }

  @override
  Future<void> get delete async {
    try {
      Response resp = await api.delete(
        '/api/products/${temp[0]}',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
    } catch (x) {
      throw 'Erro!';
    }
  }

  @override
  List<String> get headers => [
    'Código',
    'Preço',
    'Nome',
    'GTIN',
    'Unidade',
    'Descrição',
    'Classe',
    'Estoque',
    'Tributação',
  ];

  @override
  List<bool> get show => [
    true,
    true,
    true,
    true,
    true,
    true,
    true,
    true,
    true,
  ];

  @override
  Future<void> get fetch async {
    List<List> payload = [];
    try {
      Response resp = await api.get(
        '/api/products',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      for (Map<String, dynamic> j in resp.data['data']) {
        payload.add([
          j['code'],
          j['price'],
          j['name'],
          j['gtin'],
          j['um'],
          j['description'],
          j['class_id'],
          j['stock'],
          j['vtribute'],
        ]);
      }
    } catch (x) {
      throw 'Erro!';
    }
    fetched = [
      ['', '', '', '', '', '', '', '', ''],
      ...payload,
    ];
  }

  @override
  Future<void> get update async {
    try {
      Response resp = await api.put(
        '/api/products/${temp[0]}',
        data: {
          'price': temp[1],
          'name': temp[2],
          'gtin': temp[3],
          'um': temp[4],
          'description': temp[5],
          'class_id': temp[6],
          'stock': temp[7],
          'vtribute': temp[8],
        },
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
    } catch (x) {
      throw 'Erro!';
    }
  }

  @override
  List<Fld> get fields => [
    FStr().opt.ned,
    FDbl().opt.ned,
    FStr().opt.ned,
    FStr().opt.ned,
    FStr().opt.ned,
    FLongStr().opt.ned,
    FInt().opt.ned,
    FInt().opt.ned,
    FDbl().opt.ned,
  ];
}

class Pb extends Tb<P$> {
  Pb(super.source);
}