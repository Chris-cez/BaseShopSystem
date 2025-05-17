import 'package:brasil_fields/brasil_fields.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'g.dart';

abstract class ES {}

class ESN implements ES {}

class ESL implements ES {}

class ESS implements ES {}

class ESV implements ES {
  final String pjMessage;
  final String pswMessage;
  ESV({this.pjMessage = '', this.pswMessage = ''});
}

class ESA implements ES {}

abstract class EE {}

class EET implements EE {
  final String pj;
  final String psw;

  EET(this.pj, this.psw);
}

abstract class EC {
  Future<bool> login(String pj, String psw);
}

class EB extends Bloc<EE, ES> {
  final EC ec;
  EB(this.ec) : super(ESN()) {
    on<EET>((event, emit) async {
      if (event.pj.isEmpty | event.psw.isEmpty) {
        emit(
          ESV(
            pjMessage: event.pj.isEmpty ? 'Obrigatório' : '',
            pswMessage: event.psw.isEmpty ? 'Obrigatório' : '',
          ),
        );
      } else {
        emit(ESL());
        if (await ec.login(event.pj, event.psw)) {
          emit(ESS());
        } else {
          emit(ESA());
        }
      }
    });
  }
}

class EW extends StatefulWidget {
  final EC ec;

  const EW({required this.ec, super.key});

  @override
  State<EW> createState() => _EWS();
}

class _EWS extends State<EW> {
  TextEditingController pjC = TextEditingController(),
      pswC = TextEditingController();

  @override
  Widget build(BuildContext context) => Scaffold(
    body: Container(
      decoration: BoxDecoration(
        gradient: gradient
      ),
      child: Center(
        child: Container(
          constraints: BoxConstraints.tightFor(width: 250),
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
          child: BlocProvider(
            create: (context) => EB(widget.ec),
            child: BlocListener<EB, ES>(
              listener: (context, state) {
                if (state is ESS) {
                  Navigator.pushReplacementNamed(context, '/m');
                } else if (state is ESA) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(content: Text('Não foi possível autenticar')),
                  );
                }
              },
              child: BlocBuilder<EB, ES>(
                builder: (context, state) {
                  return Padding(
                    padding: const EdgeInsets.all(defaultSpacing),
                    child: Column(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        TextField(
                          enabled: state is! ESL,
                          controller: pjC,
                          inputFormatters: [
                            FilteringTextInputFormatter.digitsOnly,
                            CnpjInputFormatter(),
                          ],
                          decoration: InputDecoration(
                            border: inputBorder,
                            hintText: 'CNPJ',
                            errorText: state is ESV ? state.pjMessage : null,
                          ),
                        ),
                        space,
                        TextField(
                          enabled: state is! ESL,
                          obscureText: true,
                          controller: pswC,
                          decoration: InputDecoration(
                            border: inputBorder,
                            hintText: 'Senha',
                            errorText: state is ESV ? state.pswMessage : null,
                          ),
                        ),
                        space,
                        SizedBox(
                          width: double.infinity,
                          height: buttonHeight,
                          child: ElevatedButton(
                            style: greenButtonStyle,
                            onPressed: (state is ESL) ? (){} : () {
                              final pj = pjC.text;
                              final psw = pswC.text;
                              context.read<EB>().add(EET(pj, psw));
                            },
                            child:
                                (state is ESL)
                                    ? const SizedBox(
                                      height: 32,
                                      width: 32,
                                      child: CircularProgressIndicator(
                                        color: Colors.white,
                                      ),
                                    )
                                    : const Text('Entrar'),
                          ),
                        ),
                      ],
                    ),
                  );
                },
              ),
            ),
          ),
        ),
      ),
    ),
  );
}
