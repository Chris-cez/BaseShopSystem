import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'g.dart';

abstract class XS {}

class XSN implements XS {}

class XSL implements XS {}

class XSS implements XS {}

abstract class XE {}

class XET implements XE {}

class XB extends Bloc<XE, XS> {
  XB() : super(XSN()) {
    on<XET>((event, emit) async {
      emit(XSL());
      await key.delete;
      emit(XSS());
    });
  }
}

class XW extends StatefulWidget {
  const XW({super.key});

  @override
  State<XW> createState() => _XWS();
}

class _XWS extends State<XW> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: DecoratedBox(
        decoration: BoxDecoration(gradient: gradient),
        child: BlocProvider<XB>(
          create: (_) => XB()..add(XET()),
          child: BlocListener<XB, XS>(
            listener: (context, state) {
              if (state is XSS) {
                Navigator.of(context).pushNamedAndRemoveUntil('/e', (route) => false);
              }
            },
            child: Center(
              child: CircularProgressIndicator(color: Colors.white),
            ),
          ),
        ),
      ),
    );
  }
}
