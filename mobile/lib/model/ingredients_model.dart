class Ingredient {
  int id;
  int recipeId;
  String nome;

  Ingredient({required this.id, required this.recipeId, required this.nome});

  factory Ingredient.fromJson(Map<String, dynamic> json) {
    return Ingredient(
      id: json['Id'],
      recipeId: json['RecipeId'],
      nome: json['Nome'],
    );
  }
}
