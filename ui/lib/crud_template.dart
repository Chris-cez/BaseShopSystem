import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

abstract class Fld {
  bool optional = false;
  bool editable = true;
  Widget build(BuildContext context, List data, int index);
  Fld get opt {
    optional = false;
    return this;
  }

  Fld get ned {
    editable = false;
    return this;
  }

  Fld get ed {
    editable = true;
    return this;
  }
}

class FStr extends Fld {
  @override
  Widget build(BuildContext context, List data, int index) {
    TextEditingController controller = TextEditingController(
      text: '${data[index]}',
    );
    return TextField(
      onChanged: (s) {
        data[index] = s;
      },
      controller: controller,
      readOnly: !editable,
      decoration: InputDecoration(
        enabledBorder: InputBorder.none,
        focusedBorder: InputBorder.none,
        border: InputBorder.none,
      ),
    );
  }
}

class FLongStr extends Fld {
  @override
  Widget build(BuildContext context, List data, int index) {
    return TextButton(
      onPressed: () async {
        String? text = await showDialog(
          barrierDismissible: false,
          context: context,
          builder: (context) {
            TextEditingController controller = TextEditingController(
              text: '${data[index]}',
            );
            return AlertDialog(
              actions: [
                TextButton(
                  onPressed: () {
                    Navigator.of(
                      context,
                    ).pop(editable ? controller.text : null);
                  },
                  child: Text('OK'),
                ),
              ],
              content: TextField(
                controller: controller,
                readOnly: !editable,
                decoration: InputDecoration(
                  enabledBorder: InputBorder.none,
                  focusedBorder: InputBorder.none,
                  border: InputBorder.none,
                ),
              ),
            );
          },
        );
        if (text != null) {
          data[index] = text;
        }
      },
      child: Text(
        'Abrir',
        style: TextStyle(decoration: TextDecoration.underline),
      ),
    );
  }
}

class FInt extends Fld {
  @override
  Widget build(BuildContext context, List data, int index) {
    TextEditingController controller = TextEditingController(
      text: '${data[index]}',
    );
    return TextField(
      onChanged: (s) {
        data[index] = int.parse(s);
      },
      controller: controller,
      readOnly: !editable,
      decoration: InputDecoration(
        enabledBorder: InputBorder.none,
        focusedBorder: InputBorder.none,
        border: InputBorder.none,
      ),
    );
  }
}

class FDbl extends Fld {
  @override
  Widget build(BuildContext context, List data, int index) {
    TextEditingController controller = TextEditingController(
      text: '${data[index]}'.replaceAll(r'.', ','),
    );
    return TextField(
      onChanged: (s) {
        data[index] = double.parse(s.replaceAll(r',', '.'));
      },
      controller: controller,
      readOnly: !editable,
      decoration: InputDecoration(
        enabledBorder: InputBorder.none,
        focusedBorder: InputBorder.none,
        border: InputBorder.none,
      ),
    );
  }
}

// ignore: must_be_immutable
class T<S extends Source, B extends Tb<S>> extends StatefulWidget {
  B Function(BuildContext) bloc;

  T(this.bloc, {super.key});

  @override
  State<T<S, B>> createState() => _Ts<S, B>();
}

abstract class Source {
  int index = -1;
  int editingIndex = -1;
  List<List> fetched = [];
  List temp = [];
  Future<void> get fetch;
  Future<void> get delete;
  Future<void> get update;
  Future<void> get create;
  List<String> get headers;
  List<Fld> get fields;
  List<bool> get show;
}

abstract class TemplateEvent {}

class CreateEvent extends TemplateEvent {}

class DeleteEvent extends TemplateEvent {}

class UpdateEvent extends TemplateEvent {}

class ReadEvent extends TemplateEvent {}

abstract class CRUDState {}

class CRUDNeutral extends CRUDState {}

class CreateLoading extends CRUDState {}

class DeleteLoading extends CRUDState {}

class UpdateLoading extends CRUDState {}

class ReadLoading extends CRUDState {}

class CRUDError extends CRUDState {
  String message;
  CRUDError(this.message);
}

class CreateError extends CRUDError {
  CreateError(super.message);
}

class DeleteError extends CRUDError {
  DeleteError(super.message);
}

class UpdateError extends CRUDError {
  UpdateError(super.message);
}

class ReadError extends CRUDError {
  ReadError(super.message);
}

class CreateSuccess extends CRUDState {}

class DeleteSuccess extends CRUDState {}

class UpdateSuccess extends CRUDState {}

class ReadSuccess extends CRUDState {}

abstract class Tb<S extends Source> extends Bloc<TemplateEvent, CRUDState> {
  S source;

  Tb(this.source) : super(CRUDNeutral()) {
    on<CreateEvent>(create);
    on<UpdateEvent>(update);
    on<DeleteEvent>(delete);
    on<ReadEvent>(read);
  }
  void create(CreateEvent event, Emitter<CRUDState> emit) async {
    emit(CreateLoading());
    try {
      await source.create;
    } catch (x) {
      emit(CreateError("$x"));
      return;
    }
    emit(CreateSuccess());
  }

  void update(UpdateEvent event, Emitter<CRUDState> emit) async {
    emit(UpdateLoading());
    try {
      await source.update;
    } catch (x) {
      emit(UpdateError("$x"));
      return;
    }
    emit(UpdateSuccess());
  }

  void delete(DeleteEvent event, Emitter<CRUDState> emit) async {
    emit(DeleteLoading());
    try {
      await source.delete;
    } catch (x) {
      emit(DeleteError("$x"));
      return;
    }
    emit(DeleteSuccess());
  }

  void read(ReadEvent event, Emitter<CRUDState> emit) async {
    emit(ReadLoading());
    try {
      await source.fetch;
    } catch (x) {
      emit(ReadError("$x"));
      return;
    }
    emit(ReadSuccess());
  }
}

class _Ts<S extends Source, B extends Tb<S>> extends State<T<S, B>> {
  @override
  Widget build(BuildContext context) {
    return BlocProvider<B>(
      create: (c) => widget.bloc(c)..add(ReadEvent()),
      child: BlocListener<B, CRUDState>(
        listener: (context, state) {
          if (state is CreateSuccess) {
            context.read<B>().add(ReadEvent());
          } else if (state is DeleteSuccess) {
            context.read<B>().add(ReadEvent());
          } else if (state is UpdateSuccess) {
            context.read<B>().add(ReadEvent());
          }
        },
        child: Column(
          mainAxisSize: MainAxisSize.max,
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisAlignment: MainAxisAlignment.start,
          children: [
            BlocBuilder<B, CRUDState>(
              builder: (context, state) {
                return Row(
                  children: [
                    IconButton(
                      onPressed: () {
                        context.read<B>().add(ReadEvent());
                      },
                      icon: Icon(Icons.refresh),
                    ),
                  ],
                );
              },
            ),
            Expanded(
              child: BlocBuilder<B, CRUDState>(
                builder: (context, state) {
                  if (state is CRUDError) {
                    return Center(
                      child: Text(
                        state.message,
                        style: TextStyle(
                          color: Colors.red,
                          fontWeight: FontWeight.bold,
                          fontSize: 24,
                        ),
                      ),
                    );
                  } else if (state is ReadLoading) {
                    return Center(child: CircularProgressIndicator());
                  } else if (state is ReadSuccess) {
                    S $ = context.read<B>().source;
                    if ($.fetched.isNotEmpty) {
                      List<DataRow> rows = [];
                      for (int i = 0; i < $.fetched.length; i++) {
                        $.index = i;
                        DataCell ops = DataCell(SizedBox.shrink());
                        if (i == 0) {
                          ops = DataCell(
                            Row(
                              mainAxisAlignment: MainAxisAlignment.end,
                              crossAxisAlignment: CrossAxisAlignment.end,
                              children: [
                                IconButton(
                                  onPressed: () {
                                    context.read<B>().add(CreateEvent());
                                  },
                                  icon: Icon(Icons.save),
                                ),
                              ],
                            ),
                          );
                        } else if ($.editingIndex != $.index) {
                          ops = DataCell(
                            Row(
                              mainAxisAlignment: MainAxisAlignment.end,
                              crossAxisAlignment: CrossAxisAlignment.end,
                              children: [
                                IconButton(
                                  onPressed: () {
                                    setState(() {
                                      $.temp = [...$.fetched[i]];
                                      $.editingIndex = i;
                                    });
                                  },
                                  icon: Icon(Icons.edit),
                                ),
                                IconButton(
                                  onPressed: () {
                                    setState(() {
                                      $.index = i;
                                      $.temp = [...$.fetched[i]];
                                      context.read<B>().add(DeleteEvent());
                                    });
                                  },
                                  icon: Icon(Icons.delete),
                                ),
                              ],
                            ),
                          );
                        } else {
                          ops = DataCell(
                            Row(
                              mainAxisAlignment: MainAxisAlignment.end,
                              crossAxisAlignment: CrossAxisAlignment.end,
                              children: [
                                IconButton(
                                  onPressed: () {
                                    context.read<B>().add(UpdateEvent());
                                  },
                                  icon: Icon(Icons.save),
                                ),
                                IconButton(
                                  onPressed: () {
                                    setState(() {
                                      $.editingIndex = -1;
                                    });
                                  },
                                  icon: Icon(Icons.cancel),
                                ),
                              ],
                            ),
                          );
                        }
                        List<Fld> fls = $.fields;
                        List<DataCell> dcells = [];
                        if (i == 0) {
                          for (int X = 0; X < fls.length; X++) {
                            if ($.show[X]) {
                              dcells.add(
                                DataCell(
                                  fls[X].ed.build(context, $.fetched[0], X),
                                ),
                              );
                            }
                          }
                        } else if ($.editingIndex != $.index) {
                          for (int X = 0; X < fls.length; X++) {
                            if ($.show[X]) {
                              dcells.add(
                                DataCell(
                                  fls[X].build(context, $.fetched[i], X),
                                ),
                              );
                            }
                          }
                        } else {
                          for (int X = 0; X < fls.length; X++) {
                            if ($.show[X]) {
                              dcells.add(
                                DataCell(fls[X].ed.build(context, $.temp, X)),
                              );
                            }
                          }
                        }
                        rows.add(
                          DataRow(
                            color: WidgetStateColor.resolveWith((states) {
                              return i % 2 == 0
                                  ? Colors.grey.shade300
                                  : Colors.white;
                            }),
                            cells: [...dcells, ops],
                          ),
                        );
                      }
                      List<DataColumn> columns = [];
                      for (int X = 0; X < $.headers.length; X++) {
                        if ($.show[X]) {
                          columns.add(
                            DataColumn(
                              label: Text(
                                $.headers[X],
                                style: TextStyle(fontWeight: FontWeight.bold),
                              ),
                            ),
                          );
                        }
                      }
                      return DataTable(
                        columns: [
                          ...columns,
                          DataColumn(
                            label: Text(
                              'Operações',
                              style: TextStyle(fontWeight: FontWeight.bold),
                            ),
                            columnWidth: FlexColumnWidth(1),
                            headingRowAlignment: MainAxisAlignment.end,
                          ),
                        ],
                        rows: rows,
                      );
                    } else {
                      return Center(child: Text('Sem dados'));
                    }
                  } else if (state is ReadError) {
                    return Center(child: Text('Erro'));
                  }
                  return Text('Erro Inesperado');
                },
              ),
            ),
          ],
        ),
      ),
    );
  }
}
