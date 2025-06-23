DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tipo_uf') THEN
    CREATE TYPE tipo_uf AS ENUM (
      'AC', 'AL', 'AP', 'AM', 'BA', 'CE', 'DF', 'ES', 'GO', 'MA',
      'MT', 'MS', 'MG', 'PA', 'PB', 'PR', 'PE', 'PI', 'RJ', 'RN',
      'RS', 'RO', 'RR', 'SC', 'SP', 'SE', 'TO'
    );
  END IF;

  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'raca_cor') THEN
    CREATE TYPE raca_cor AS ENUM (
      'Branca', 'Preta', 'Parda', 'Indigena', 'Outros'
    );
  END IF;

  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'escolaridade') THEN
    CREATE TYPE escolaridade AS ENUM (
      'Analfabeta',
      'Ensino Fundamental Incompleto',
      'Ensino Fundamental Completo',
      'Ensino Médio Completo',
      'Ensino Superior Completo'
    );
  END IF;

  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'resposta_simples') THEN
    CREATE TYPE resposta_simples AS ENUM ('Sim', 'Não', 'Não sabe');
  END IF;

  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'resposta_corrimento') THEN
    CREATE TYPE resposta_corrimento AS ENUM ('Sim', 'Não', 'Não sabe', 'Não lembra');
  END IF;

  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'resposta_sangramento') THEN
    CREATE TYPE resposta_sangramento AS ENUM ('Sim', 'Não', 'Não sabe', 'Não lembra', 'Não está em menopausa');
  END IF;

  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'resposta_menopausa') THEN
    CREATE TYPE resposta_menopausa AS ENUM ('Sim', 'Não', 'Não lembra');
  END IF;

  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tipo_exame') THEN
    CREATE TYPE tipo_exame AS ENUM ('Rastreamento', 'Seguimento');
  END IF;

  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_exame') THEN
    CREATE TYPE status_exame AS ENUM ('Aprovado', 'Rejeitado', 'Pendente');
  END IF;

  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tipo_usuario') THEN
    CREATE TYPE tipo_usuario AS ENUM ('paciente', 'profissional');
  END IF;
END$$;

-- Tabelas

CREATE TABLE IF NOT EXISTS municipio (
  cod CHAR(7) PRIMARY KEY,
  nome VARCHAR(100) NOT NULL,
  uf tipo_uf NOT NULL
);

CREATE TABLE IF NOT EXISTS paciente_infos (
  id SERIAL PRIMARY KEY,
  cartao_sus VARCHAR(15) NOT NULL,
  cpf_paciente CHAR(11) NOT NULL UNIQUE,
  nome_completo VARCHAR(150) NOT NULL,
  data_nascimento DATE NOT NULL,

  cep CHAR(8) NOT NULL,
  ddd CHAR(2) NOT NULL,
  telefone CHAR(9) NOT NULL,
  nacionalidade VARCHAR(100) NOT NULL,
  uf tipo_uf NOT NULL,

  raca_cor raca_cor,
  escolaridade escolaridade,
  nome_mae VARCHAR(150),
  nome_social VARCHAR(150),
  logradouro VARCHAR(255),
  numero_residencia VARCHAR(20),
  complemento VARCHAR(100),
  setor VARCHAR(100),
  cod_municipio CHAR(7) REFERENCES municipio(cod),
  ponto_referencia VARCHAR(255),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT un_cartao_sus_cpf UNIQUE (cartao_sus, cpf_paciente)
);

CREATE TABLE IF NOT EXISTS ubs_infos (
  id SERIAL PRIMARY KEY,
  cnes VARCHAR NOT NULL,
  unidade VARCHAR(255) NOT NULL,
  municipios_ubs TEXT,
  estado tipo_uf,
  uf tipo_uf NOT NULL,
  protocolo VARCHAR NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS profissional_saude (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(150) NOT NULL,
  cpf CHAR(11) NOT NULL UNIQUE,
  email VARCHAR(150),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS exame_citopatologico (
  id SERIAL PRIMARY KEY,
  id_profissional INTEGER REFERENCES profissional_saude(id),
  paciente_id INTEGER REFERENCES paciente_infos(id),

  motivo_exame tipo_exame NOT NULL,
  primeira_vez_exame resposta_simples NOT NULL,
  usa_diu resposta_simples NOT NULL,
  usa_anticoncepcional resposta_simples NOT NULL,
  esta_gestante resposta_simples NOT NULL,
  usa_hormonio resposta_simples NOT NULL,
  ja_fez_radioterapia resposta_simples NOT NULL,

  data_ultima_menstruacao DATE,
  esta_menopausa resposta_menopausa NOT NULL,
  teve_corrimento resposta_corrimento NOT NULL,
  teve_sangramento_anormal resposta_sangramento NOT NULL,

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS resultado_exame (
  id SERIAL PRIMARY KEY,
  exame_id INTEGER REFERENCES exame_citopatologico(id) ON DELETE CASCADE,

  status status_exame NOT NULL,
  resultado VARCHAR(255),
  motivo_rejeicao TEXT,
  observacoes TEXT,
  data_avaliacao DATE DEFAULT CURRENT_DATE,
  profissional_responsavel VARCHAR(255),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS usuario (
  id SERIAL PRIMARY KEY,
  tipo tipo_usuario NOT NULL,
  ref_id INTEGER NOT NULL,  
  CPF CHAR(11) NOT NULL UNIQUE,
  senha TEXT NOT NULL
);

-- Garantir que a constraint seja atualizada corretamente
ALTER TABLE usuario DROP CONSTRAINT IF EXISTS senha;

ALTER TABLE usuario ADD CONSTRAINT senha CHECK (
  senha ~ '[A-Z]'
  AND senha ~ '[a-z]'
  AND senha ~ '[0-9]'
  AND senha ~ '[^A-Za-z0-9]'
);
