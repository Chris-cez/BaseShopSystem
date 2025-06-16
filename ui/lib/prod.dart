import 'package:login_bss/crud_template.dart';

class P$ extends Source {

  @override
  Future<void> get create async {
    return;
  }

  @override
  Future<void> get delete async {
    return;
  }

  @override
  List<String> get headers => [
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
    true,
    true,
    true,
    true,
    true,
    true,
    true,
    true,
    true,
  ];

  @override
  Future<void> get fetch async {
    fetched = [['', '', '', '', '', '', '', '', '']];
  }

  @override
  Future<void> get update async {
    return;
  }

  @override
  List<Fld> get fields => [
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