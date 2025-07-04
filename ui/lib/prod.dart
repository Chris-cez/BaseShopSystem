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
          'code': fetched[0][1],
          'price': fetched[0][2],
          'name': fetched[0][3],
          'gtin': fetched[0][4], // temp[3],
          'um': fetched[0][5],
          'description': fetched[0][6], // temp[5],
          'class_id': fetched[0][7], //fetched[0][6],
          'stock': fetched[0][8], //temp[7],
          'valtrib': fetched[0][9] // temp[8],
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
    } on DioException catch (a, x) {
      throw 'Erro! ${a.response!.data}';
    }
  }

  @override
  List<String> get headers => [
    'ID',
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
    false,
    true,
    true,
    true,
    true,
    true,
    true,
    false,
    true,
    false,
  ];

  @override
  Future<void> get fetch async {
    List<List> payload = [];
    try {
      Response resp = await api.get(
        '/api/products',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      print(resp.data);
      for (Map<String, dynamic> j in resp.data['data']) {
        payload.add([
          j['ID'],
          j['code'],
          j['price'],
          j['name'],
          j['gtin'],
          j['um'],
          j['description'],
          j['class_id'],
          j['stock'],
          j['valtrib'],
        ]);
      }
    } catch (x) {
      throw 'Erro!';
    }
    fetched = [
      ['', '', '', '', '', '', '', 1, '', 1],
      ...payload,
    ];
    print(fetched);
  }

  @override
  Future<void> get update async {
    try {
      Response resp = await api.put(
        '/api/products/${temp[0]}',
        data: {
          'code': temp[1],
          'price': temp[2],
          'name': temp[3],
          'gtin': temp[4],
          'um': temp[5],
          'description': temp[6],
          'class_id': temp[7],
          'stock': temp[8],
          'valtrib': temp[9],
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