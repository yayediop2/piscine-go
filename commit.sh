#!/bin/bash

# Ajouter les dépôts distants
 git remote add origin https://learn.zone01dakar.sn/git/yayediop/piscine-go
 git remote add github https://github.com/yayediop2/piscine-go

# Ajouter tous les fichiers au commit
git add .

# Faire un commit avec un message spécifié par l'utilisateur
echo "Entrez votre message de commit: "
read commit_message
git commit -m "$commit_message"

# Pousser les changements au dépôt par défaut (origin)
git push

# Renommer la branche master en main
git branch -m master main

# Pousser la branche main sur GitHub
git push github main 

# Renommer la branche master localement 
git branch -m master
