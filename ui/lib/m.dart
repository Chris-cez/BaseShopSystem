import 'package:flutter/material.dart';
import 'package:login_bss/cli.dart';
import 'package:login_bss/crud_template.dart';
import 'package:login_bss/inv.dart';
import 'package:login_bss/prod.dart';
import 'package:login_bss/vendas.dart'; 
import 'g.dart';

class MW extends StatefulWidget {
  const MW({super.key});

  @override
  State<MW> createState() => _MWS();
}

final List<String> pageNames = ['Início', 'Produtos', 'Vendas', 'Clientes'];

class _MWS extends State<MW> {
  int pageIndex = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: DecoratedBox(
        decoration: BoxDecoration(gradient: gradient),
        child: Padding(
          padding: EdgeInsets.all(defaultSpacing),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.start,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(children: [Spacer()]),
              Container(
                padding: EdgeInsets.all(16),
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(16),
                  border: Border.all(color: Colors.grey, width: 1),
                  boxShadow: const [
                    BoxShadow(
                      color: Colors.black12,
                      blurRadius: 8,
                      offset: Offset(0, 4),
                    ),
                  ],
                ),
                child: Wrap(
                  runSpacing: defaultSpacing,
                  spacing: defaultSpacing,
                  children: [
                    SizedBox(
                      height: buttonHeight,
                      child: ElevatedButton(
                        style: redButtonStyle,
                        onPressed: () async {
                          Navigator.of(
                            context,
                          ).pushNamedAndRemoveUntil('/x', (route) => false);
                        },
                        child: const Text('Sair'),
                      ),
                    ),
                    ...pageNames.map((name) {
                      return SizedBox(
                        height: buttonHeight,
                        child: ElevatedButton(
                          style: blueButtonStyle,
                          onPressed: () {
                            setState(() {
                              pageIndex = pageNames.indexOf(name);
                            });
                          },
                          child: Text(name),
                        ),
                      );
                    }),
                  ],
                ),
              ),
              space,
              Expanded(
                child: Container(
                  decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.circular(16),
                    border: Border.all(color: Colors.grey, width: 1),
                    boxShadow: const [
                      BoxShadow(
                        color: Colors.black12,
                        blurRadius: 8,
                        offset: Offset(0, 4),
                      ),
                    ],
                  ),
                  child: Column(
                    children: [
                      Row(children: [Spacer()]),
                      Expanded(child: _buildContent(context)),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildContent(BuildContext context) {
    if (pageIndex == 1) {
      return T<P$, Pb>((context) => Pb(P$()));
    } else if (pageIndex == 2) {
      return T<I$, Ib>((context) => Ib(I$()))..mode="C";
    } else if (pageIndex == 3) {
      return T<C$, Cb>((context) => Cb(C$()));
    }
    return Text('Página inicial');
  }
}
