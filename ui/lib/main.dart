import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:login_bss/e.dart';
import 'package:login_bss/g.dart';
import 'package:login_bss/m.dart';
import 'package:login_bss/x.dart';

void main() {
  runApp(const MyApp());
}

class ECx extends EC {
  @override
  Future<bool> login(String pj, String psw) async {
    try {
      Response response = await api.post(
        '/api/entrar',
        data: {'cnpj': pj, 'password': psw},
      );
      if (response.statusCode == 200) {
        key.value = response.data['key'];
        //key.expiration = DateTime.now().add(const Duration(days: 1));
        await key.store;
        return true;
      }
    } catch (_) {
      return false;
    }
    return false;
  }
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'BSS',
      home: EW(ec: ECx()),
      routes: {
        '/e': (_) => EW(ec: ECx()),
        '/m': (_) => MW(),
        '/x': (_) => XW(),
      },
    );
  }
}
