import 'dart:convert';
import 'package:http/http.dart' as http;

import '../model/ingredients_model.dart';
import '../model/recipe_model.dart';

class ApiService {
  final String baseUrl = "http://localhost:8888";

  Future<List<Recipe>> getAllRecipes() async {
    final response = await http.get(Uri.parse('$baseUrl/receita'));
    if (response.statusCode == 200) {
      List<dynamic> data = json.decode(response.body)['data'];
      return data.map((recipe) => Recipe.fromJson(recipe)).toList();
    } else {
      throw Exception('erro');
    }
  }

  Future<int> addRecipe(Recipe recipe) async {
    final response = await http.post(
      Uri.parse('$baseUrl/receita'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({
        'nome': recipe.nome,
        'tempo_preparo': recipe.tempoPreparo,
        'custo_aproximado': recipe.custoAproximado,
      }),
    );

    if (response.statusCode == 200 || response.statusCode == 201) {
      final Map<String, dynamic> responseBody = json.decode(response.body);
      final int recipeId = responseBody['data']['Id'];
      return recipeId;
    } else {
      throw Exception('erro');
    }
  }

  Future<void> updateRecipe(Recipe recipe) async {
    final response = await http.put(
      Uri.parse('$baseUrl/receita'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({
        'id': recipe.id,
        'nome': recipe.nome,
        'tempo_preparo': recipe.tempoPreparo,
        'custo_aproximado': recipe.custoAproximado,
      }),
    );
    if (response.statusCode != 200) {
      throw Exception('erro');
    }
  }

  Future<void> deleteRecipe(int recipeId) async {
    final response = await http.delete(Uri.parse('$baseUrl/receita/id/$recipeId'));
    if (response.statusCode != 200) {
      throw Exception('erro');
    }
  }

  Future<List<Ingredient>> getIngredientsByRecipeId(int recipeId) async {
    final response = await http.get(Uri.parse('$baseUrl/ingredients/id/$recipeId'));
    if (response.statusCode == 200) {
      List<dynamic> data = json.decode(response.body)['data'];
      return data.map((ingredient) => Ingredient.fromJson(ingredient)).toList();
    } else {
      throw Exception('erro');
    }
  }

  Future<void> addIngredient(Ingredient ingredient) async {
    final response = await http.post(
      Uri.parse('$baseUrl/ingredients'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({
        'recipe_id': ingredient.recipeId,
        'nome': ingredient.nome,
      }),
    );
    if (response.statusCode != 200) {
      throw Exception('erro');
    }
  }

  Future<void> deleteIngredient(int ingredientId) async {
    final response = await http.delete(Uri.parse('$baseUrl/ingredients/id/$ingredientId'));
    if (response.statusCode != 200) {
      throw Exception('erro');
    }
  }
}
