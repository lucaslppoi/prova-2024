import 'package:flutter/material.dart';

import '../model/ingredients_model.dart';
import '../model/recipe_model.dart';
import '../service/api_service.dart';

class RecipeDetailPage extends StatefulWidget {
  final Recipe recipe;

  const RecipeDetailPage({super.key, required this.recipe});

  @override
  _RecipeDetailPageState createState() => _RecipeDetailPageState();
}

class _RecipeDetailPageState extends State<RecipeDetailPage> {
  late Future<List<Ingredient>> _ingredients;
  final ApiService _apiService = ApiService();

  final TextEditingController _nomeController = TextEditingController();
  final TextEditingController _tempoPreparoController = TextEditingController();
  final TextEditingController _custoAproximadoController = TextEditingController();

  @override
  void initState() {
    super.initState();

    _nomeController.text = widget.recipe.nome;
    _tempoPreparoController.text = widget.recipe.tempoPreparo.toString();
    _custoAproximadoController.text = widget.recipe.custoAproximado.toStringAsFixed(2);

    _ingredients = _apiService.getIngredientsByRecipeId(widget.recipe.id);
  }

  void _saveRecipe() async {
    final String nome = _nomeController.text;
    final int tempoPreparo = int.tryParse(_tempoPreparoController.text) ?? 0;
    final double custoAproximado = double.tryParse(_custoAproximadoController.text) ?? 0.0;

      final updatedRecipe = Recipe(
        id: widget.recipe.id,
        nome: nome,
        tempoPreparo: tempoPreparo,
        custoAproximado: custoAproximado,
      );

      await _apiService.updateRecipe(updatedRecipe);
  }

  void _addIngredient() async {
      final newIngredient = Ingredient(
        id: 0,
        recipeId: widget.recipe.id,
        nome: _ingredientController.text,
      );

      await _apiService.addIngredient(newIngredient);

      setState(() {
        _ingredients = _apiService.getIngredientsByRecipeId(widget.recipe.id);
      });

      _ingredientController.clear();
  }

  final TextEditingController _ingredientController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.recipe.nome),
        actions: [
          IconButton(
            icon: const Icon(Icons.delete),
            onPressed: () {
              _apiService.deleteRecipe(widget.recipe.id).then((_) {
                Navigator.pop(context);
              });
            },
          )
        ],
      ),
      body: Column(
        children: [
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                TextField(
                  controller: _nomeController,
                  decoration: const InputDecoration(labelText: 'nome'),
                ),
                const SizedBox(height: 8),
                TextField(
                  controller: _tempoPreparoController,
                  keyboardType: TextInputType.number,
                  decoration: const InputDecoration(labelText: 'tempo de preparo'),
                ),
                const SizedBox(height: 8),
                TextField(
                  controller: _custoAproximadoController,
                  keyboardType: TextInputType.number,
                  decoration: const InputDecoration(labelText: 'custo aproximado'),
                ),
                const SizedBox(height: 16),
                ElevatedButton(
                  onPressed: _saveRecipe,
                  child: const Text('salvar alteracoes'),
                ),
                const Divider(height: 30, thickness: 2),
              ],
            ),
          ),
          Expanded(
            child: FutureBuilder<List<Ingredient>>(
              future: _ingredients,
              builder: (context, snapshot) {
                  return ListView.builder(
                    itemCount: snapshot.data!.length,
                    itemBuilder: (context, index) {
                      final ingredient = snapshot.data![index];
                      return ListTile(
                        title: Text(ingredient.nome),
                        trailing: Row(
                          mainAxisSize: MainAxisSize.min,
                          children: [
                            IconButton(
                              icon: const Icon(Icons.delete),
                              onPressed: () {
                                _apiService
                                    .deleteIngredient(ingredient.id)
                                    .then((_) {
                                  setState(() {
                                    _ingredients = _apiService
                                        .getIngredientsByRecipeId(widget.recipe.id);
                                  });
                                });
                              },
                            ),
                          ],
                        ),
                      );
                    },
                  );
              },
            ),
          ),
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: Row(
              children: [
                Expanded(
                  child: TextField(
                    controller: _ingredientController,
                    decoration: const InputDecoration(labelText: 'adicionar ingrediente'),
                  ),
                ),
                IconButton(
                  icon: const Icon(Icons.add),
                  onPressed: _addIngredient,
                )
              ],
            ),
          ),
        ],
      ),
    );
  }
}