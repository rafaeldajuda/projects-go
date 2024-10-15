# Exercício 1: Nível Fácil – Leitura de Arquivo com Tratamento de Erro

**Descrição:**

Crie um programa que tente abrir um arquivo de texto e leia seu conteúdo. Caso o arquivo não exista ou ocorra um erro durante a leitura, o programa deve exibir uma mensagem de erro apropriada. Utilize o pacote os para abrir o arquivo e tratar o erro de forma adequada.

**Dicas:**

* Use os.Open() para abrir o arquivo.
* Utilize defer para fechar o arquivo após a leitura.
* Trate o erro retornado por os.Open() para verificar se o arquivo existe.