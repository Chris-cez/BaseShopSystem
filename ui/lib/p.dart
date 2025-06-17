import 'package:flutter/material.dart';
import 'g.dart';

class PW extends StatefulWidget {
  const PW({super.key});

  @override
  State<PW> createState() => _PWS();
}

class _PWS extends State<PW> {
  final TextEditingController nameController = TextEditingController();
  final TextEditingController priceController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        const Text(
          'Cadastro de Produto',
          style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
        ),
        space,
        TextField(
          controller: nameController,
          decoration: const InputDecoration(
            labelText: 'Placeholder',
            border: inputBorder,
          ),
        ),
        space,
        TextField(
          controller: priceController,
          keyboardType: TextInputType.number,
          decoration: const InputDecoration(
            labelText: 'Placeholder',
            border: inputBorder,
          ),
        ),
        space,
        SizedBox(
          height: buttonHeight,
          width: double.infinity,
          child: ElevatedButton(
            style: greenButtonStyle,
            onPressed: () {
            },
            child: const Text('Placeholder'),
          ),
        ),
      ],
    );
  }
}
