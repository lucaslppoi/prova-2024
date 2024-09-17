import 'package:flutter/material.dart';
import 'package:mobile/view/recipe_detail_page.dart';

import '../model/recipe_model.dart';
import '../service/api_service.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  late Future<List<Recipe>> _recipes;
  final ApiService _apiService = ApiService();

  final TextEditingController _nomeController = TextEditingController();
  final TextEditingController _tempoPreparoController = TextEditingController();
  final TextEditingController _custoAproximadoController = TextEditingController();

  @override
  void initState() {
    super.initState();
    _recipes = _apiService.getAllRecipes();
  }

  void _addRecipe() async {
    final String nome = _nomeController.text;
    final int tempoPreparo = int.tryParse(_tempoPreparoController.text) ?? 0;
    final double custoAproximado = double.tryParse(_custoAproximadoController.text) ?? 0.0;

      final newRecipe = Recipe(
        id: 0,
        nome: nome,
        tempoPreparo: tempoPreparo,
        custoAproximado: custoAproximado,
      );
      await _apiService.addRecipe(newRecipe);

      setState(() {
        _recipes = _apiService.getAllRecipes();
      });

      _nomeController.clear();
      _tempoPreparoController.clear();
      _custoAproximadoController.clear();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('receitas'),
      ),
      body: Row(
        children: [
          Expanded(
            child: FutureBuilder<List<Recipe>>(
              future: _recipes,
              builder: (context, snapshot) {
                  return ListView.builder(
                    itemCount: snapshot.data!.length,
                    itemBuilder: (context, index) {
                      return ListTile(
                        title: Text(snapshot.data![index].nome),
                        onTap: () {
                          Navigator.push(
                            context,
                            MaterialPageRoute(
                              builder: (context) => RecipeDetailPage(
                                recipe: snapshot.data![index],
                              ),
                            ),
                          );
                        },
                      );
                    },
                  );
              },
            ),
          ),
          Expanded(
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const Text(
                    'criar receita',
                  ),
                  TextField(
                    controller: _nomeController,
                    decoration: const InputDecoration(labelText: 'nome'),
                  ),
                  TextField(
                    controller: _tempoPreparoController,
                    keyboardType: TextInputType.number,
                    decoration:
                    const InputDecoration(labelText: 'tempo de preparo'),
                  ),
                  TextField(
                    controller: _custoAproximadoController,
                    keyboardType: TextInputType.number,
                    decoration:
                    const InputDecoration(labelText: 'custo aproximado'),
                  ),
                  const SizedBox(height: 20),
                  ElevatedButton(
                    onPressed: _addRecipe,
                    child: const Text('adicionar receita'),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}