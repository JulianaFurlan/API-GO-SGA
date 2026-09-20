#!/bin/bash

BASE="http://localhost:8080/api/v1"

echo "== health =="
curl -s "$BASE/health"; echo

echo "== criar sala s1 (capacidade 1) =="
curl -s -X POST "$BASE/salas" -H "Content-Type: application/json" \
  -d '{"id":"s1","nome":"Sala 101","capacidade":1,"recursos":["projetor"]}'; echo

echo "== criar aluno a1 =="
curl -s -X POST "$BASE/alunos" -H "Content-Type: application/json" \
  -d '{"matricula":"a1","nome":"Fulano","email_institucional":"fulano@edu.unifil.br"}'; echo

echo "== criar aluno a2 =="
curl -s -X POST "$BASE/alunos" -H "Content-Type: application/json" \
  -d '{"matricula":"a2","nome":"Ciclano","email_institucional":"ciclano@edu.unifil.br"}'; echo

echo "== criar turma t1 =="
curl -s -X POST "$BASE/turmas" -H "Content-Type: application/json" \
  -d '{"id":"t1","nome":"Turma A","disciplina":"Go","docente":"Prof X"}'; echo

echo "== matricular a1 em t1 (espera 201) =="
curl -s -w " [status=%{http_code}]\n" -X POST "$BASE/turmas/t1/alunos" -H "Content-Type: application/json" \
  -d '{"aluno_id":"a1"}'

echo "== matricular a1 de novo (espera 409, duplicado) =="
curl -s -w " [status=%{http_code}]\n" -X POST "$BASE/turmas/t1/alunos" -H "Content-Type: application/json" \
  -d '{"aluno_id":"a1"}'

echo "== listar turmas (espera quantidade_alunos=1, alocada=false) =="
curl -s "$BASE/turmas"; echo

echo "== alocar sala s1 na turma t1 (espera 201) =="
curl -s -w " [status=%{http_code}]\n" -X POST "$BASE/turmas/t1/alocar" -H "Content-Type: application/json" \
  -d '{"sala_id":"s1","dia_semana":"segunda","horario_inicio":"08:00","horario_termino":"10:00"}'

echo "== matricular a2 em t1 (sala lotada, espera 422) =="
curl -s -w " [status=%{http_code}]\n" -X POST "$BASE/turmas/t1/alunos" -H "Content-Type: application/json" \
  -d '{"aluno_id":"a2"}'

echo "== agenda da sala s1 =="
curl -s "$BASE/salas/s1/agenda"; echo

echo "== alunos da turma t1 =="
curl -s "$BASE/turmas/t1/alunos"; echo

echo "== matricular em turma inexistente (espera 404) =="
curl -s -w " [status=%{http_code}]\n" -X POST "$BASE/turmas/naoexiste/alunos" -H "Content-Type: application/json" \
  -d '{"aluno_id":"a1"}'

echo "== FIM =="
