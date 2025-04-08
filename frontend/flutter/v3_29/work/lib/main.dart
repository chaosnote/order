import 'package:flutter/material.dart';
import 'dart:developer' as developer ;

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: '登入表單',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      initialRoute: '/',
      routes: {
        '/': (context) => LoginForm(),
        '/register': (context) => RegisterPage(), // 註冊頁面路由
      },
    );
  }
}

class LoginForm extends StatefulWidget {
  const LoginForm({super.key});

  @override
  LoginFormState createState() => LoginFormState();
}

class LoginFormState extends State<LoginForm> {
  LoginFormState();

  final _formKey = GlobalKey<FormState>();
  final _usernameController = TextEditingController();
  final _passwordController = TextEditingController();

  @override
  void dispose() {
    _usernameController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold( // 使用 Scaffold 提供基本的應用程式結構
      appBar: AppBar(title: Text('登入')),
      body: Center( // 將 body 包在 Center Widget 中
        child: SizedBox(
          width: 300,
          child: Padding(
            padding: const EdgeInsets.all(16.0),
            child: Form(
              key: _formKey,
              child: Column(
                mainAxisSize: MainAxisSize.min, // 讓 Column 佔據最小的空間
                crossAxisAlignment: CrossAxisAlignment.stretch, // 讓 Column 的子元件填滿寬度
                children: <Widget>[
                  TextFormField(
                    controller: _usernameController,
                    decoration: InputDecoration(labelText: '使用者名稱：'),
                    validator: (value) {
                      if (value == null || value.isEmpty) {
                        return '請輸入使用者名稱';
                      }
                      // 正則表達式驗證
                      final regex = RegExp(r'^[a-zA-Z][a-zA-Z0-9]{3,7}$'); // 4 ~ 8
                      if (!regex.hasMatch(value)) {
                        return '使用者名稱格式不正確';
                      }
                      return null;
                    },
                  ),
                  TextFormField(
                    controller: _passwordController,
                    decoration: InputDecoration(labelText: '密碼：'),
                    obscureText: true,
                    validator: (value) {
                      if (value == null || value.isEmpty) {
                        return '請輸入密碼';
                      }
                      // 正則表達式驗證
                      final regex = RegExp(r'^[a-zA-Z0-9]{3,7}$'); // 4 ~ 8
                      if (!regex.hasMatch(value)) {
                        return '密碼格式不正確';
                      }
                      return null;
                    },
                  ),
                  SizedBox(height: 16.0),
                  Padding(
                    padding: const EdgeInsets.symmetric(vertical: 16.0),
                    child: ElevatedButton(
                      onPressed: () {
                        if (_formKey.currentState!.validate()) {
                          developer.log('使用者名稱: ${_usernameController.text}');
                          developer.log('密碼: ${_passwordController.text}');
                          // 在此處新增您的登入邏輯，例如發送 HTTP 請求
                        }
                      },
                      child: Text('登入'),
                    ),
                  ),
                  TextButton(
                    onPressed: () {
                      Navigator.pushNamed(context, '/register');
                    },
                    child: Text('註冊'),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}

class RegisterPage extends StatelessWidget { // 簡單的註冊頁面
  const RegisterPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('註冊')),
      body: Center(
        child: Text('註冊頁面'),
      ),
    );
  }
}