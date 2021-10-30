import 'package:exam_flutter/screens/done_screen.dart';
import 'package:exam_flutter/screens/todo_screen.dart';
import 'package:flutter/material.dart';

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);
  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final _tab =<Tab>[
    Tab(text: "未処理",icon: Icon(Icons.remove_red_eye)),
    Tab(text: "処理",icon: Icon(Icons.radar_rounded)),
  ];

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(length: _tab.length, child: Scaffold(
      appBar: AppBar(
        title: const Text("マムシ"),
        bottom: TabBar(
          tabs: _tab,
        ),
      ),
      body: TabBarView(
        children: <Widget>[
          DoneTodo(),
          Todo(),
        ],
      ),
    ));
  }
}
