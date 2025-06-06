import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:login_bss/g.dart';

abstract class PLS {}

class PLSN implements PLS {}

class PLSL implements PLS {}

class PLSS implements PLS {
  List<dynamic> l = [];
  PLSS(this.l);

}

abstract class PLE {}

class PLET implements PLE {}

abstract class PLC {
  Future<dynamic> getProducts();
}

abstract class EC {
  Future<bool> login(String pj, String psw);
}

class PLB extends Bloc<PLE, PLS> {
  final PLC plc;
  PLB(this.plc) : super(PLSN()) {
    on<PLET>((event, emit) async {
      emit(PLSL());
      dynamic products = await plc.getProducts();
      emit(PLSS(products));
    });
  }
}

class PLW extends StatefulWidget {
  const PLW({super.key});

  @override
  State<PLW> createState() => PLWS();
}

class GetProducts extends PLC {
  @override
  Future getProducts() async {
    Response response = await api.get('/api/products');
    print(response.data);
    return response.data['data'];
  }
}

class PLWS extends State<PLW> {
  @override
  Widget build(BuildContext context) => BlocProvider(
    create: (_) => PLB(GetProducts())..add(PLET()),
    child: BlocConsumer<PLB, PLS>(
      builder: (context, state) {
        if (state is PLSL) {
          return Text("Carregando");
        } else if (state is PLSS) {
          return Text("${state.l}");
        }
        return Text('???');
      },
      listener: (BuildContext context, PLS state) {},
    ),
  );
}
