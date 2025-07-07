import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:brasil_fields/brasil_fields.dart';
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
      if (event.pj.isEmpty || event.psw.isEmpty) {
        emit(ESV(
          pjMessage: event.pj.isEmpty ? 'Obrigatório' : '',
          pswMessage: event.psw.isEmpty ? 'Obrigatório' : '',
        ));
        return;
      }
      emit(ESL());
      final success = await ec.login(event.pj, event.psw);
      emit(success ? ESS() : ESA());
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
  final pjC = TextEditingController();
  final pswC = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        decoration: BoxDecoration(gradient: gradient),
        alignment: Alignment.center,
        child: BlocProvider(
          create: (_) => EB(widget.ec),
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
                final loading = state is ESL;
                final errorState = state is ESV ? state : null;

                return Container(
                  width: 320,
                  padding: const EdgeInsets.all(24),
                  decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.circular(16),
                    border: Border.all(color: Colors.grey.shade300),
                    boxShadow: const [
                      BoxShadow(
                        color: Colors.black12,
                        blurRadius: 10,
                        offset: Offset(0, 6),
                      ),
                    ],
                  ),
                  child: Column(
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      const Text(
                        'Login',
                        style: TextStyle(
                          fontSize: 24,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                      const SizedBox(height: 16),
                      TextField(
                        controller: pjC,
                        enabled: !loading,
                        inputFormatters: [
                          FilteringTextInputFormatter.digitsOnly,
                          CnpjInputFormatter(),
                        ],
                        decoration: InputDecoration(
                          labelText: 'CNPJ',
                          border: inputBorder,
                          errorText: errorState?.pjMessage,
                        ),
                      ),
                      const SizedBox(height: 16),
                      TextField(
                        controller: pswC,
                        enabled: !loading,
                        obscureText: true,
                        decoration: InputDecoration(
                          labelText: 'Senha',
                          border: inputBorder,
                          errorText: errorState?.pswMessage,
                        ),
                      ),
                      const SizedBox(height: 24),
                      SizedBox(
                        width: double.infinity,
                        height: buttonHeight,
                        child: ElevatedButton(
                          style: greenButtonStyle,
                          onPressed: loading
                              ? null
                              : () {
                                  final pj = pjC.text;
                                  final psw = pswC.text;
                                  context.read<EB>().add(EET(pj, psw));
                                },
                          child: loading
                              ? const CircularProgressIndicator(
                                  color: Colors.white, strokeWidth: 2)
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
    );
  }
}
