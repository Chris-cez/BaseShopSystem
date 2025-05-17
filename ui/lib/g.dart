import 'package:dio/dio.dart';
import 'package:flutter/material.dart';

class Key {
  String? value;
  DateTime? expiration;
  Key({this.value, this.expiration});
  bool get valid {
    if (expiration != null) {
      return expiration!.isBefore(DateTime.now()) && value != null;
    }
    return value != null;
  }

  Future<void> get store async {
    await Future.delayed(const Duration(seconds: 1));
    print('Chave armazenada: $value');
    print('Chave armazenada temporariamente: $value');
  }

  Future<void> get delete async {
    value = null;
    await Future.delayed(const Duration(seconds: 1));
    print('Chave deletada: $value');
  }

  Future<void> get load async {
    await Future.delayed(const Duration(seconds: 1));
    print('Chave carregada: $value');
  }
}

Key key = Key();

Dio api = Dio(BaseOptions(
  baseUrl: 'http://localhost:8080',
  connectTimeout: const Duration(seconds: 5),
  receiveTimeout: const Duration(seconds: 5),
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },
  responseType: ResponseType.json,
));

final buttonShape = RoundedRectangleBorder(borderRadius: BorderRadius.circular(8));

final redButtonStyle = ElevatedButton.styleFrom(
  foregroundColor: Colors.white,
  backgroundColor: Colors.red,
  shape: buttonShape,
);

final blueButtonStyle = ElevatedButton.styleFrom(
  foregroundColor: Colors.white,
  backgroundColor: Colors.blue,
  shape: buttonShape,
);

final greenButtonStyle = ElevatedButton.styleFrom(
  foregroundColor: Colors.white,
  backgroundColor: Colors.green,
  shape: buttonShape,
);

final gradient = LinearGradient(
  colors: [Colors.red, Colors.purple, Colors.green],
  begin: Alignment.topLeft,
  end: Alignment.bottomRight,
);

const inputBorder = OutlineInputBorder(
  borderRadius: BorderRadius.all(Radius.circular(8)),
  borderSide: BorderSide(color: Colors.grey, width: 1),
);

const buttonHeight = 48.0;

const defaultSpacing = 16.0;
final space = const SizedBox(width: 16, height: 16);

