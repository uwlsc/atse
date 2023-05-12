import 'package:flutter/material.dart';
import 'package:google_nav_bar/google_nav_bar.dart';

class Homepage extends StatefulWidget{
  const Homepage({Key? key}) : super(key:key);


  @override
  State<Homepage> createState() => _HomePageState();
}

class _HomePageState extends State<Homepage>{
  int _selectIndex = 0;

  void _navigateBottomBar(int index){
    setState(() {
      _selectIndex = index;
    });
  }
  final List<Widget> _pages = [
    Center(
        child: Text(
          'PDF CONVERTER',
          style: TextStyle(fontSize: 50),
        )
      ),
      Center(
        child: Text(
          'Search',
          style: TextStyle(fontSize: 50),
        )
      ),
      Center(
        child: Text(
          'Login',
          style: TextStyle(fontSize: 50),
        )
      ),
  ];
  @override
  Widget build(BuildContext context){
    var gNav = GNav(
       selectedIndex: _selectIndex,
        
        
        backgroundColor: Colors.black,
        color: Colors.white,
        activeColor: Colors.white ,
        tabBackgroundColor: Colors.grey.shade800,
        onTabChange: _navigateBottomBar,
        tabs: [
          GButton(
            icon: Icons.home,
          text: 'Home'),
          GButton(icon: Icons.search,
          text: 'Search'
          ),
          GButton(icon: Icons.login,
          text: 'Login'
          ),
        ],
      );
    return Scaffold(
      body: _pages[_selectIndex],
      bottomNavigationBar: gNav,
    );
  }
}