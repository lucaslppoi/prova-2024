class Recipe {
  int id;
  String nome;
  int tempoPreparo;
  double custoAproximado;

  Recipe({
    required this.id,
    required this.nome,
    required this.tempoPreparo,
    required this.custoAproximado,
  });

  factory Recipe.fromJson(Map<String, dynamic> json) {
    return Recipe(
      id: json['Id'],
      nome: json['Nome'],
      tempoPreparo: json['TempoPreparo'],
      custoAproximado: json['CustoAproximado'],
    );
  }
}
