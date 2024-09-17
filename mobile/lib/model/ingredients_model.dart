class IngredientModel {
  String nome;
  int recipeId;
  int id;

  IngredientModel({
    required this.nome,
    required this.id,
    required this.recipeId,
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'nome': nome,
      'id': id,
      'recipeId': recipeId,
    };
  }

  factory IngredientModel.fromJson(Map<String, dynamic> json) {
    return IngredientModel(
      nome: json['nome'] as String,
      id: json['id'] as int,
      recipeId: json['recipeId'] as int,
    );
  }
}
