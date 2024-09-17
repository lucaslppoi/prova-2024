class RecipeModel {
  String nome;
  int id;
  int tempoPreparo;
  double custoAproximado;

  RecipeModel({
    required this.nome,
    required this.id,
    required this.tempoPreparo,
    required this.custoAproximado,
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'nome': nome,
      'id': id,
      'tempoPreparo': tempoPreparo,
      'custoAproximado': custoAproximado,
    };
  }

  factory RecipeModel.fromJson(Map<String, dynamic> json) {
    return RecipeModel(
      nome: json['nome'] as String,
      id: json['id'] as int,
      tempoPreparo: json['tempoPreparo'] as int,
      custoAproximado: json['custoAproximado'] as double,
    );
  }
}
