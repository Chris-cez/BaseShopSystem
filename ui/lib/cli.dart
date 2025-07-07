import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:login_bss/crud_template.dart';
import 'package:login_bss/g.dart';

class C$ extends Source {

  @override
  Future<void> get create async {
    print(fetched[0]);
    try {
      Response resp = await api.post(
        '/api/clients',
        data: {
          'cpf': fetched[0][1],
          'name': fetched[0][2],
          'address_id': fetched[0][3],
        },
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
    } on DioException catch(a,x) {
      throw 'Erro! ${a.response!.data}';
    }
  }

  @override
  Future<void> get delete async {
    try {
      Response resp = await api.delete(
        '/api/clients/${temp[0]}',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
    } on DioException catch (a, x) {
      throw 'Erro! ${a.response!.data}';
    }
  }

  @override
  List<String> get headers => [
    'ID',
    'CPF',
    'Nome',
    'Endereço',
  ];

  @override
  List<bool> get show => [
    false,
    true,
    true,
    false,
  ];

  @override
  Future<void> get fetch async {
    List<List> payload = [];
    try {
      Response resp = await api.get(
        '/api/clients',
        options: Options(headers: {'Authorization': 'Bearer ${key.value}'}),
      );
      print(resp.data);
      for (Map<String, dynamic> j in resp.data['data']) {
        payload.add([
          j['ID'],
          j['cpf'],
          j['name'],
          j['address_id'],
        ]);
      }
    } catch (x) {
      throw 'Erro!';
    }
    fetched = [
      ['', '', '', 1],
      ...payload,
    ];
    print(fetched);
  }

  @override
  Future<void> get update async {
    try {
      Response resp = await api.put(
        '/api/clients/${temp[0]}',
        data: {
          'cpf': temp[1],
          'name': temp[2],
          'address_id': temp[3],
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
    FStr().opt.ned,
    FInt().opt.ned,
  ];
}

class Cb extends Tb<C$> {
  Cb(super.source);
}

class FCli extends Fld {
  @override
  Widget build(BuildContext context, List data, int index) {
    Widget w = T<C$, Cb>((context) => Cb(C$()))
    ..mode="SELECT"
    ..isSelected=(d){
      return d[0] == data[index];
    };
    return TextButton(onPressed: editable ? () async {
      List? cli = await showDialog(context: context, builder: (context)=>AlertDialog(content: w,));
      if (cli != null) {
        data[index] = cli[0];
      }
    } : (){}, child: Text("Abrir"));
  }
}