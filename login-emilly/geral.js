

(function() { 
    'use strict';
    console.log("geral.js: IIFE (Função Auto-Executável) iniciada.");


    const dadosPacienteAutenticado = {
        nomeCompleto: "Maria Linda da Silva",
        primeiroNome: "Maria",
        dataNascimento: "01/01/1985",
        endereco: {
            cep: "12345-678",
            cidade: "Exemplópolis",
            logradouro: "Rua dos bobos, n°0"
        },
        cpf: "012.345.678-90",
        cartaoSUS: "123.4567.7890.1234",
        online: {
            email: "maria88@linda.com.br",
            login: "maria88@linda.com.br",
            senha: "Pi2025"
        },
        rg: "00.000.000-0",
        sexo: "Feminino",
        celular: "(11) 98765-4321",
        telefone: "(11) 1234-5678",
        escolaridade: "Superior Completo",
        observacao: "Tem alzheimer e uma cirurgia no estômago."
        // fotoUrl: "caminho/para/foto_real_do_paciente.jpg" 
    };
    console.log("geral.js: 'dadosPacienteAutenticado' definido:", dadosPacienteAutenticado);

    // --------------------------------------------------------------------
    // FUNÇÃO PARA POPULAR OS DADOS DO PACIENTE NO HTML
    // --------------------------------------------------------------------
    function popularDadosPaciente() {
        console.log("geral.js: Função 'popularDadosPaciente' foi chamada.");
        try {
            const pacienteFotoPrincipalEl = document.getElementById('pacienteFotoPrincipal');
            if (pacienteFotoPrincipalEl) {
                if (dadosPacienteAutenticado.fotoUrl) {
                    pacienteFotoPrincipalEl.src = dadosPacienteAutenticado.fotoUrl;
                } else {
                    console.log("geral.js: 'dadosPacienteAutenticado.fotoUrl' não encontrado. Usando imagem placeholder do HTML.");
                }
            } else {
                console.warn("geral.js: Elemento HTML com ID 'pacienteFotoPrincipal' não encontrado.");
            }

            function setTextContent(id, value) {
                const element = document.getElementById(id);
                if (element) {
                    element.textContent = value;
                } else {
                    console.warn(`geral.js: Elemento HTML com ID '${id}' não encontrado.`);
                }
            }

            setTextContent('pacienteNomePrincipal', dadosPacienteAutenticado.nomeCompleto);
            setTextContent('pacienteNomeCompleto', dadosPacienteAutenticado.nomeCompleto);
            setTextContent('pacienteDataNascimento', dadosPacienteAutenticado.dataNascimento);
            setTextContent('pacienteCPF', dadosPacienteAutenticado.cpf);
            setTextContent('pacienteRG', dadosPacienteAutenticado.rg);
            setTextContent('pacienteSexo', dadosPacienteAutenticado.sexo);
            setTextContent('pacienteCartaoSUS', dadosPacienteAutenticado.cartaoSUS);
            
            const enderecoCompleto = `${dadosPacienteAutenticado.endereco.logradouro}, ${dadosPacienteAutenticado.endereco.cidade}`;
            setTextContent('pacienteEnderecoCompleto', enderecoCompleto);
            setTextContent('pacienteCEP', dadosPacienteAutenticado.endereco.cep);
            
            setTextContent('pacienteCelular', dadosPacienteAutenticado.celular);
            setTextContent('pacienteTelefone', dadosPacienteAutenticado.telefone);
            setTextContent('pacienteEmail', dadosPacienteAutenticado.online.email);
            setTextContent('pacienteEscolaridade', dadosPacienteAutenticado.escolaridade);
            // setTextContent('pacienteObservacao', dadosPacienteAutenticado.observacao); // Seu HTML combinado não tem 'pacienteObservacao' na seção de perfil, mas sim em um local diferente. Verifique.

            console.log("geral.js: Função 'popularDadosPaciente' concluída.");
        } catch (error) {
            console.error("geral.js: Erro DENTRO da função 'popularDadosPaciente':", error);
        }
    }

    // --------------------------------------------------------------------
    // FUNÇÃO PARA CONFIGURAR AS ABAS (Perfil / Histórico) - da sua versão original de perfil.html
    // ATENÇÃO: Se você está usando a página combinada 'paciente_detalhes.html' com os botões
    // 'btnMostrarPerfil' e 'btnMostrarHistorico' e a lógica de abas que forneci para ela,
    // esta função 'setupTabsPaciente' pode não ser mais necessária ou precisará ser adaptada,
    // pois ela procura por IDs como 'btnProntuario' e 'btnFichaCitopatologica'.
    // --------------------------------------------------------------------
    function setupTabsPaciente() {
        console.log("geral.js: Função 'setupTabsPaciente' foi chamada.");
        const btnProntuario = document.getElementById('btnProntuario'); // ID da sua perfil.html original
        const btnFichaCitopatologica = document.getElementById('btnFichaCitopatologica'); // ID da sua perfil.html original
        const perfilHeaderInfo = document.getElementById('perfilHeaderInfo');
        const perfilPacienteContainer = document.getElementById('perfilPacienteContainer'); // ID da sua perfil.html original
        const historicoHospitalarContainer = document.getElementById('historicoHospitalarContainer'); // ID da sua perfil.html original

        if (!btnProntuario || !btnFichaCitopatologica || !perfilHeaderInfo || !perfilPacienteContainer || !historicoHospitalarContainer) {
            console.warn("geral.js: Um ou mais elementos HTML para as abas (setupTabsPaciente) não foram encontrados. A funcionalidade de abas pode estar comprometida se esta função for necessária.");
            return;
        }
        
        function mostrarPerfil() {
            perfilHeaderInfo.style.display = ''; 
            perfilPacienteContainer.style.display = ''; 
            historicoHospitalarContainer.style.display = 'none';
            btnProntuario.classList.add('ativo');
            btnFichaCitopatologica.classList.remove('ativo');
            console.log("geral.js: Aba 'Perfil' mostrada por setupTabsPaciente.");
        }

        function mostrarHistorico() {
            // NOTA: Na sua lógica original de setupTabsPaciente, você estava escondendo 'perfilHeaderInfo'.
            // Na página combinada, o 'perfilHeaderInfo' (identificação do paciente) geralmente fica visível acima das abas.
            // Ajuste conforme o comportamento desejado.
            perfilHeaderInfo.style.display = 'none'; // Considere se quer manter isso.
            perfilPacienteContainer.style.display = 'none';
            historicoHospitalarContainer.style.display = ''; 
            btnProntuario.classList.remove('ativo');
            btnFichaCitopatologica.classList.add('ativo');
            console.log("geral.js: Aba 'Histórico' mostrada por setupTabsPaciente.");
        }

        btnProntuario.addEventListener('click', mostrarPerfil);
        btnFichaCitopatologica.addEventListener('click', mostrarHistorico);

        // Estado inicial baseado na classe 'ativo' dos botões (se existirem)
        if (btnProntuario.classList.contains('ativo')) {
            mostrarPerfil();
        } else if (btnFichaCitopatologica.classList.contains('ativo')) {
            mostrarHistorico();
        } else {
            mostrarPerfil(); // Padrão
        }
        console.log("geral.js: Função 'setupTabsPaciente' concluída.");
    }
    

    // --------------------------------------------------------------------
    // DADOS DA INSTITUIÇÃO (mantidos)
    // --------------------------------------------------------------------
    const dadosInstituicaoAutenticada = {
        nomeInstituicao: "Joana", 
        cnes: "1234567",
        crmCorenResponsavel: "CRM/SP 123456",
        cnpj: "12.345.678/0001-99", 
        email: "contato@clinicajoana.com", 
        celular: "(11) 91234-5678", 
        telefone: "(11) 5555-1234", 
        cep: "01000-001", 
        cidade: "São Paulo", 
        endereco: "Rua da Clínica, Bairro Saúde", 
        numero: "100", 
        complemento: "Andar 5, Sala 502", 
        bairro: "Saúde", 
        online: {
            senha: "Pi2025"
        }
    };

    // --------------------------------------------------------------------
    // FUNÇÕES GLOBAIS (Podem ser usadas em qualquer lugar)
    // --------------------------------------------------------------------
    function calcularIdade(dataNascimentoString) {
        if (!dataNascimentoString || !/^\d{2}\/\d{2}\/\d{4}$/.test(dataNascimentoString)) { return "N/D"; }
        const partesData = dataNascimentoString.split('/');
        const diaNasc = parseInt(partesData[0], 10);
        const mesNasc = parseInt(partesData[1], 10) - 1;
        const anoNasc = parseInt(partesData[2], 10);
        const dataNascDate = new Date(anoNasc, mesNasc, diaNasc);
        if (isNaN(dataNascDate.getTime())) return "N/D";
        const hoje = new Date();
        let idade = hoje.getFullYear() - dataNascDate.getFullYear();
        const m = hoje.getMonth() - dataNascDate.getMonth();
        if (m < 0 || (m === 0 && hoje.getDate() < dataNascDate.getDate())) { idade--; }
        return idade >= 0 ? idade : "N/D";
    }

    // --------------------------------------------------------------------
    // FUNÇÃO PARA PREENCHER E GERENCIAR O FORMULÁRIO DE "ALTERAR DADOS" (PACIENTE)
    // --------------------------------------------------------------------
   function preencherCamposFormularioEdicaoPaciente() {
    console.log("geral.js: Tentando preencher campos do formulário de edição...");
    if (typeof dadosPacienteAutenticado === 'undefined') {
        console.error("geral.js: Objeto 'dadosPacienteAutenticado' não definido internamente.");
        return;
    }

    // Função auxiliar para definir o valor de um campo de input
    function setInputValue(id, value) {
        const element = document.getElementById(id);
        if (element) {
            element.value = value !== undefined && value !== null ? value : '';
        } else {
            // Comentado para não poluir o console se o ID não for desta página
            // console.warn(`geral.js: Elemento input com ID '${id}' não encontrado para preenchimento.`);
        }
    }

    // Função auxiliar para definir o texto de um elemento (span, p, etc.)
    function setTextContent(id, value) {
        const element = document.getElementById(id);
        if (element) {
            element.textContent = value !== undefined && value !== null ? value : '';
        } else {
            // Comentado para não poluir o console se o ID não for desta página
            // console.warn(`geral.js: Elemento com ID '${id}' não encontrado para preenchimento.`);
        }
    }
    
    // Popular saudação e nome de exibição no perfil
    setTextContent('saudacaoNomePaciente', dadosPacienteAutenticado.primeiroNome || (dadosPacienteAutenticado.nomeCompleto ? dadosPacienteAutenticado.nomeCompleto.split(" ")[0] : ''));
    setTextContent('nomeDisplayPerfil', dadosPacienteAutenticado.nomeCompleto);

    // Popular campos do formulário
    setInputValue('inputNome', dadosPacienteAutenticado.nomeCompleto);
    setInputValue('inputEmail', dadosPacienteAutenticado.online ? dadosPacienteAutenticado.online.email : '');
    setInputValue('inputCPF', dadosPacienteAutenticado.cpf);
    setInputValue('inputCelular', dadosPacienteAutenticado.celular);
    setInputValue('inputTelefone', dadosPacienteAutenticado.telefone);
    setInputValue('inputEscolaridade', dadosPacienteAutenticado.escolaridade);
    
    if (dadosPacienteAutenticado.endereco) {
        setInputValue('inputCEP', dadosPacienteAutenticado.endereco.cep);
        setInputValue('inputCidade', dadosPacienteAutenticado.endereco.cidade);
        setInputValue('inputEndereco', dadosPacienteAutenticado.endereco.logradouro);
    } else {
        console.warn("geral.js: Dados de endereço não encontrados em 'dadosPacienteAutenticado'.");
    }

    // Popular avatar se fotoUrl existir
    const avatarContainer = document.getElementById('avatarContainer');
    if (avatarContainer) {
        if (dadosPacienteAutenticado.fotoUrl) { // fotoUrl não está no seu objeto de exemplo atual
            avatarContainer.innerHTML = `<img src="${dadosPacienteAutenticado.fotoUrl}" alt="Foto de ${dadosPacienteAutenticado.primeiroNome || 'paciente'}">`;
        } else {
            // console.log("geral.js: Nenhuma fotoUrl definida em dadosPacienteAutenticado para o avatar.");
        }
    }
    console.log("geral.js: Campos do formulário de edição deveriam estar populados.");
}

// Expor a função para ser chamada globalmente
// Coloque isso antes do })(); final da IIFE
window.preencherFormularioEdicaoGlobal = preencherCamposFormularioEdicaoPaciente; 
console.log("geral.js: Função 'preencherFormularioEdicaoGlobal' exposta.");

// })(); // Esta é a linha final da sua IIFE

    // --------------------------------------------------------------------
    // FUNÇÃO PARA PREENCHER E GERENCIAR O FORMULÁRIO DE "ALTERAR DADOS" (INSTITUIÇÃO)
    // --------------------------------------------------------------------
    function preencherFormularioInstituicaoInterna() {
        // ... (seu código original para preencherFormularioInstituicaoInterna) ...
        const nomeInst = dadosInstituicaoAutenticada.nomeInstituicao || "Nome da Instituição";
        const saudacaoNomeEl = document.getElementById('saudacao-nome');
        if (saudacaoNomeEl) {
            saudacaoNomeEl.textContent = nomeInst;
        }
        const nomeExibicaoPerfilEl = document.getElementById('nome-exibicao-perfil');
        if (nomeExibicaoPerfilEl) {
            nomeExibicaoPerfilEl.textContent = nomeInst;
        }
        const idsDosCamposParaPreencher = [
            'nomeInstituicao', 'cnes', 'crmCorenResponsavel', 'cnpj',
            'email', 'celular', 'telefone', 'cep', 'cidade',
            'endereco', 'numero', 'complemento', 'bairro'
        ];
        idsDosCamposParaPreencher.forEach(id => {
            const el = document.getElementById(id);
            if (el) {
                el.value = dadosInstituicaoAutenticada[id] || '';
            }
        });
        const camposDoFormulario = document.querySelectorAll('.info-formulario input[type="text"], .info-formulario input[type="email"], .info-formulario input[type="tel"]');
        const botaoEditar = document.querySelector('.info-formulario .botao.editar');
        const botaoSalvar = document.querySelector('.info-formulario .botao.salvar');

        if (botaoEditar && botaoSalvar && camposDoFormulario.length > 0) {
            camposDoFormulario.forEach(input => input.disabled = true);
            botaoEditar.style.display = 'inline-block';
            botaoSalvar.style.display = 'none';
            botaoEditar.onclick = function() {
                camposDoFormulario.forEach(input => input.disabled = false);
                botaoEditar.style.display = 'none';
                botaoSalvar.style.display = 'inline-block';
                camposDoFormulario[0].focus();
            };
            botaoSalvar.onclick = function(event) {
                event.preventDefault();
                idsDosCamposParaPreencher.forEach(id => {
                    const inputElement = document.getElementById(id);
                    if (inputElement && dadosInstituicaoAutenticada.hasOwnProperty(id)) {
                        dadosInstituicaoAutenticada[id] = inputElement.value;
                    } else if (inputElement) {
                        dadosInstituicaoAutenticada[id] = inputElement.value;
                    }
                });
                alert('Dados da instituição salvos!');
                window.location.href = "iniciomedico.html"; // NOTA: Verifique se este é o redirecionamento correto
                camposDoFormulario.forEach(input => input.disabled = true);
                botaoSalvar.style.display = 'none';
                botaoEditar.style.display = 'inline-block';
            };
        } else {
            if (!botaoEditar) console.error("Botão Editar não encontrado para o formulário da instituição.");
            if (!botaoSalvar) console.error("Botão Salvar não encontrado para o formulário da instituição.");
            if (camposDoFormulario.length === 0) console.error("Nenhum campo de formulário encontrado para a instituição.");
        }
        const fotoPerfilElement = document.querySelector('.cartao .perfil .avatar img');
        if (fotoPerfilElement) {
            fotoPerfilElement.style.cursor = 'pointer';
            fotoPerfilElement.title = 'Clique para alterar a foto';
            fotoPerfilElement.addEventListener('click', function() {
                const desejaAlterar = window.confirm("Deseja alterar a foto de perfil da instituição/profissional?");
                if (desejaAlterar) {
                    alert("Funcionalidade para alterar a foto acionada.\n(Em uma implementação real, aqui abriria um seletor de arquivos.)");
                }
            });
        }
    }
    window.preencherFormularioInstituicao = preencherFormularioInstituicaoInterna;

    // --------------------------------------------------------------------
    // EXECUÇÃO QUANDO O DOM ESTIVER PRONTO
    // --------------------------------------------------------------------
    document.addEventListener('DOMContentLoaded', () => {
        console.log("geral.js: DOMContentLoaded disparado.");

        // Chamada para popularDadosPaciente se estiver na página correta
        // NOTA: Ajuste o ID do body ou a condição se necessário para sua página combinada
        if (document.body.id === 'paginaUnificadaPaciente' || document.body.id === 'paginaPerfilPaciente' || document.getElementById('pacienteNomePrincipal')) {
            if (typeof popularDadosPaciente === 'function') {
                popularDadosPaciente();
            }
        }
        
        // Chamada para setupTabsPaciente (da sua versão original de perfil.html)
        // NOTA: Se a sua página combinada ('paciente_detalhes.html') já tem sua própria lógica de abas
        // (com botões 'btnMostrarPerfil' e 'btnMostrarHistorico'), esta função 'setupTabsPaciente'
        // pode ser redundante ou causar conflitos. Verifique se ela ainda é necessária.
        // Ela procura por IDs como 'btnProntuario' e 'btnFichaCitopatologica'.
        if (document.getElementById('btnProntuario') && document.getElementById('btnFichaCitopatologica')) { // Verifica se os elementos que setupTabsPaciente espera existem
            if (typeof setupTabsPaciente === 'function') {
                // setupTabsPaciente(); // DESCOMENTE SE VOCÊ AINDA PRECISA DESTA LÓGICA ESPECÍFICA
                console.log("geral.js: Chamada para setupTabsPaciente seria feita aqui, mas está comentada. Verifique a necessidade.");
            }
        }


        // --------------------------------------------------------------------
        // LÓGICA DA BARRA DE PESQUISA (Como estava no seu código original)
        // --------------------------------------------------------------------
        const inputPesquisa = document.querySelector('.search-bar__input');
        const botaoPesquisa = document.querySelector('.search-bar__button');

        if (inputPesquisa && botaoPesquisa) { // Adicionada verificação para segurança
            botaoPesquisa.addEventListener('click', () => {
                const termo = inputPesquisa.value.trim().toLowerCase();
                if (termo === "") {
                    alert("Digite um nome ou cartão SUS.");
                    return;
                }
                const nome = dadosPacienteAutenticado.nomeCompleto.toLowerCase();
                const sus = dadosPacienteAutenticado.cartaoSUS.replace(/\D/g, '');
                if (nome.includes(termo) || sus.includes(termo.replace(/\D/g, ''))) {
                    alert(`Paciente encontrado:\nNome: ${dadosPacienteAutenticado.nomeCompleto}\nCartão SUS: ${dadosPacienteAutenticado.cartaoSUS}`);
                    window.location.href = 'perfil.html';
                } else {
                    alert("Paciente não encontrado.");
                }
            });
        } else {
            console.warn("geral.js: Elementos da barra de pesquisa não encontrados.");
        }

        // --------------------------------------------------------------------
        // LÓGICA DO DROPDOWN DO PERFIL E MODAL DE LOGOUT (REESTRUTURADA)
        // --------------------------------------------------------------------
        const headerIconePerfilTrigger = document.getElementById('headerIconePerfilTrigger');
        const perfilDropdownOverlay = document.getElementById('perfilDropdownOverlay');
        const fecharDropdownBtn = document.getElementById('fecharDropdownBtn');
        const dropdownNomeCompleto = document.getElementById('dropdownNomeCompleto');
        const dropdownDataNascimento = document.getElementById('dropdownDataNascimento');
        const dropdownIdade = document.getElementById('dropdownIdade');
        const dropdownCEP = document.getElementById('dropdownCEP');
        const dropdownCPF = document.getElementById('dropdownCPF');
        const dropdownCartaoSUS = document.getElementById('dropdownCartaoSUS');

        const linkSairDaConta = document.getElementById('linkSairDaConta');
        const confirmLogoutModalElement = document.getElementById('confirmLogoutModal');
        const confirmLogoutModalButtonElement = document.getElementById('confirmLogoutModalButton');
        const cancelLogoutModalButtonElement = document.getElementById('cancelLogoutModalButton');

        if (!headerIconePerfilTrigger || !perfilDropdownOverlay) {
            console.warn("geral.js: Elementos essenciais do dropdown de perfil (headerIconePerfilTrigger ou perfilDropdownOverlay) não encontrados. A funcionalidade do dropdown pode estar comprometida.");
        } else {
            console.log("geral.js: Elementos do dropdown de perfil encontrados. Configurando listeners...");

            function togglePerfilDropdown() {
                if (!dadosPacienteAutenticado) {
                    console.warn("geral.js: dadosPacienteAutenticado não disponível para o dropdown.");
                    return;
                }
                
                const isVisivel = perfilDropdownOverlay.classList.contains('visivel');
                if (isVisivel) {
                    perfilDropdownOverlay.classList.remove('visivel');
                    console.log("geral.js: Dropdown de perfil ocultado.");
                } else {
                    console.log("geral.js: Preenchendo e mostrando dropdown de perfil.");
                    if (dropdownNomeCompleto) dropdownNomeCompleto.textContent = dadosPacienteAutenticado.nomeCompleto || "Nome não disponível";
                    if (dropdownDataNascimento) dropdownDataNascimento.textContent = dadosPacienteAutenticado.dataNascimento || "N/D";
                    if (dropdownIdade) {
                        const idade = calcularIdade(dadosPacienteAutenticado.dataNascimento); // calcularIdade deve estar definida globalmente na IIFE
                        dropdownIdade.textContent = idade + (idade !== "N/D" ? " anos" : "");
                    }
                    if (dropdownCEP && dadosPacienteAutenticado.endereco) dropdownCEP.textContent = dadosPacienteAutenticado.endereco.cep || "N/D";
                    else if (dropdownCEP) dropdownCEP.textContent = "N/D";
                    if (dropdownCPF) dropdownCPF.textContent = dadosPacienteAutenticado.cpf || "N/D";
                    if (dropdownCartaoSUS) dropdownCartaoSUS.textContent = dadosPacienteAutenticado.cartaoSUS || "N/D";
                    
                    perfilDropdownOverlay.classList.add('visivel'); // Usa a classe 'visivel'
                    console.log("geral.js: Dropdown de perfil tornado visível.");
                }
            }

            headerIconePerfilTrigger.addEventListener('click', function(event) {
                event.preventDefault(); 
                event.stopPropagation(); 
                console.log("geral.js: Ícone do perfil (headerIconePerfilTrigger) clicado.");
                togglePerfilDropdown();
            });

            if (fecharDropdownBtn) {
                fecharDropdownBtn.addEventListener('click', function(event) { 
                    event.stopPropagation(); 
                    console.log("geral.js: Botão fechar dropdown (fecharDropdownBtn) clicado.");
                    perfilDropdownOverlay.classList.remove('visivel');
                });
            } else {
                console.warn("geral.js: Botão para fechar dropdown (fecharDropdownBtn) não encontrado.");
            }

            document.addEventListener('click', function(event) {
                if (perfilDropdownOverlay.classList.contains('visivel')) {
                    if (!perfilDropdownOverlay.contains(event.target) && !headerIconePerfilTrigger.contains(event.target)) {
                        console.log("geral.js: Clique fora do dropdown. Fechando.");
                        perfilDropdownOverlay.classList.remove('visivel');
                    }
                }
            });
            
            document.addEventListener('keydown', function(event) {
                if (event.key === 'Escape' && perfilDropdownOverlay.classList.contains('visivel')) {
                    console.log("geral.js: Tecla ESC pressionada. Fechando dropdown.");
                    perfilDropdownOverlay.classList.remove('visivel');
                }
            });

            if (linkSairDaConta && confirmLogoutModalElement && confirmLogoutModalButtonElement && cancelLogoutModalButtonElement) {
                linkSairDaConta.addEventListener('click', function(event) {
                    event.preventDefault(); 
                    if (perfilDropdownOverlay.classList.contains('visivel')) {
                        perfilDropdownOverlay.classList.remove('visivel'); 
                    }
                    confirmLogoutModalElement.classList.add('visivel'); 
                    console.log("geral.js: Link 'Sair da Conta' clicado, modal de logout aberto.");
                });
                cancelLogoutModalButtonElement.addEventListener('click', function() {
                    confirmLogoutModalElement.classList.remove('visivel'); 
                });
                confirmLogoutModalButtonElement.addEventListener('click', function() {
                    localStorage.removeItem('pacienteLogadoIdentificador'); 
                    window.location.href = "escolher-tipo-usuario.html"; 
                });
                confirmLogoutModalElement.addEventListener('click', function(event) {
                    if (event.target === confirmLogoutModalElement) {
                        confirmLogoutModalElement.classList.remove('visivel');
                    }
                });
                document.addEventListener('keydown', function(event) {
                    if (event.key === 'Escape' && confirmLogoutModalElement.classList.contains('visivel')) {
                        confirmLogoutModalElement.classList.remove('visivel');
                    }
                });
                console.log("geral.js: Lógica do modal de logout configurada.");
            } else {
                console.warn("geral.js: Um ou mais elementos do modal de logout não encontrados.");
            }
        }

        // --------------------------------------------------------------------
        // ANTIGA "LÓGICA DO DROPDOWN DO MÉDICO" - COMENTADA PARA EVITAR CONFLITO
        // Esta seção estava no seu código original dentro do DOMContentLoaded e usava a classe 'ativo'.
        // Se os IDs ('headerIconePerfilTrigger', 'perfilDropdownOverlay') forem os mesmos
        // da lógica de dropdown acima (que usa 'visivel'), esta seção causará conflitos.
        // Se for para um dropdown *diferente*, certifique-se que os IDs são únicos.
        // --------------------------------------------------------------------
        /*
        const medicoHeaderIconePerfilTrigger = document.getElementById('headerIconePerfilTrigger'); // Mesmo ID da lógica acima
        const medicoPerfilDropdownOverlay = document.getElementById('perfilDropdownOverlay');   // Mesmo ID da lógica acima
        const medicoFecharDropdownBtn = document.getElementById('fecharDropdownBtn');         // Mesmo ID da lógica acima

        if (medicoHeaderIconePerfilTrigger && medicoPerfilDropdownOverlay && medicoFecharDropdownBtn) {
            medicoHeaderIconePerfilTrigger.addEventListener('click', function(event) {
                event.stopPropagation(); 
                // medicoPerfilDropdownOverlay.classList.toggle('ativo'); // CONFLITO: usa 'ativo', a lógica acima usa 'visivel'
                console.log("geral.js: ANTIGA LÓGICA do dropdown do médico (com ID headerIconePerfilTrigger e classe 'ativo') FOI ACIONADA. ESTÁ COMENTADA PARA EVITAR CONFLITOS.");
            });
            // ... (restante da lógica antiga que usava 'ativo') ...
        }
        console.warn("geral.js: A seção 'LÓGICA DO DROPDOWN DO MÉDICO' que usava a classe 'ativo' está comentada. Verifique se ela é necessária ou se os IDs precisam ser diferentes da lógica principal do dropdown de perfil.");
        */

        // Chamadas para preencher formulários de edição, se estiver na página correta
        // NOTA: Ajuste os IDs do body ou as condições se necessário
        if (document.getElementById('titulo-nome') && document.getElementById('nome-perfil') && typeof window.preencherFormulario === 'function') {
            window.preencherFormulario();
        }
        if (document.getElementById('saudacao-nome') && document.getElementById('nome-exibicao-perfil') && document.getElementById('nomeInstituicao') && typeof window.preencherFormularioInstituicao === 'function') {
            window.preencherFormularioInstituicao();
        }
        
        console.log("geral.js: Configurações do DOMContentLoaded finalizadas.");
    }); // Fim do DOMContentLoaded

    // --------------------------------------------------------------------
    // BLOCOS DE CÓDIGO PARA PÁGINAS ESPECÍFICAS (mantidos como no seu original)
    // --------------------------------------------------------------------
    // === Bloco para cadastro.html ===
    (function() {
        if (document.getElementById('crm') && document.querySelector('form[id="cadastroForm"]')) {
            // Código específico para cadastro.html (se houver)
        }
    })();

    // === Bloco para codigoemail.html ===
    (function() {
        const codigoInput = document.getElementById('codigoInput');
        const codigoError = document.getElementById('codigoError');
        const codigoSuccessMessage = document.getElementById('codigoSuccessMessage');
        const reenviarCodigoBtn = document.getElementById('reenviarCodigoBtn');
        const countdownSpan = document.getElementById('countdown');
        const formVerificarCodigo = document.getElementById('form-verificar-codigo');

        if (formVerificarCodigo && codigoInput && reenviarCodigoBtn && countdownSpan) {
            let countdownInterval;
            let tempoRestante = 10;
            function iniciarContadorReenvio() { /* ... seu código ... */ 
                tempoRestante = 10;
                reenviarCodigoBtn.disabled = true;
                countdownSpan.textContent = tempoRestante;
                if (countdownInterval) clearInterval(countdownInterval);
                countdownInterval = setInterval(() => {
                    tempoRestante--;
                    countdownSpan.textContent = tempoRestante;
                    if (tempoRestante <= 0) {
                        clearInterval(countdownInterval);
                        reenviarCodigoBtn.disabled = false;
                        countdownSpan.textContent = '';
                    }
                }, 1000);
            }
            iniciarContadorReenvio();
            reenviarCodigoBtn.addEventListener('click', function() {
                alert('Código reenviado!');
                iniciarContadorReenvio();
            });
            formVerificarCodigo.addEventListener('submit', function(event) { /* ... seu código ... */ 
                event.preventDefault();
                if(codigoError) codigoError.textContent = '';
                if(codigoSuccessMessage) codigoSuccessMessage.style.display = 'none';
                const codigoDigitado = codigoInput.value.trim();
                const codigoCorreto = "123456"; // Simulação
                if (codigoDigitado === "") {
                    if(codigoError) codigoError.textContent = "Por favor, insira o código.";
                    if(codigoError && codigoError.style) codigoError.style.display = 'block';
                    codigoInput.focus(); return;
                }
                if (codigoDigitado === codigoCorreto) {
                    if(codigoSuccessMessage) {
                        codigoSuccessMessage.textContent = 'Código válido! Redirecionando...';
                        codigoSuccessMessage.style.display = 'block';
                    }
                    if(codigoError && codigoError.style) codigoError.style.display = 'none';
                    setTimeout(function() { window.location.href = 'escolher-nova-senha.html'; }, 1500);
                } else {
                    if(codigoError) codigoError.textContent = "Código incorreto. Tente novamente.";
                    if(codigoError && codigoError.style) codigoError.style.display = 'block';
                    codigoInput.focus();
                }
            });
        }
    })();

    // === Bloco para escolher-nova-senha.html ===
    (function() {
        const formElement = document.getElementById('novaSenhaForm');
        const novaSenhaInputElement = document.getElementById('novaSenhaInput');
        if (formElement && novaSenhaInputElement) { /* ... seu código ... */ 
            const confirmaSenhaInputElement = document.getElementById('confirmaSenhaInput');
            const novaSenhaErrorElement = document.getElementById('novaSenhaError');
            const confirmaSenhaErrorElement = document.getElementById('confirmaSenhaError');
            formElement.addEventListener('submit', function(event) {
                event.preventDefault();
                if (novaSenhaErrorElement) { novaSenhaErrorElement.textContent = ''; novaSenhaErrorElement.style.display = 'none'; }
                if (confirmaSenhaErrorElement) { confirmaSenhaErrorElement.textContent = ''; confirmaSenhaErrorElement.style.display = 'none'; }
                const novaSenhaValue = novaSenhaInputElement.value;
                const confirmaSenhaValue = confirmaSenhaInputElement ? confirmaSenhaInputElement.value : '';
                let isFormValid = true; const MIN_SENHA_LENGTH = 6;
                if (novaSenhaValue === '') {
                    if (novaSenhaErrorElement) { novaSenhaErrorElement.textContent = 'Por favor, insira a nova senha.'; novaSenhaErrorElement.style.display = 'block'; } 
                    else { alert("Por favor, insira a nova senha."); }
                    isFormValid = false;
                } else if (novaSenhaValue.length < MIN_SENHA_LENGTH) {
                    if (novaSenhaErrorElement) { novaSenhaErrorElement.textContent = `A nova senha deve ter no mínimo ${MIN_SENHA_LENGTH} caracteres.`; novaSenhaErrorElement.style.display = 'block'; } 
                    else { alert(`A nova senha deve ter no mínimo ${MIN_SENHA_LENGTH} caracteres.`); }
                    isFormValid = false;
                }
                if (novaSenhaValue !== '' && confirmaSenhaValue === '') {
                    if (confirmaSenhaErrorElement) { confirmaSenhaErrorElement.textContent = 'Por favor, confirme a nova senha.'; confirmaSenhaErrorElement.style.display = 'block'; } 
                    else { alert("Por favor, confirme a nova senha."); }
                    isFormValid = false; 
                }
                if (novaSenhaValue !== '' && novaSenhaValue.length >= MIN_SENHA_LENGTH && confirmaSenhaValue !== '') {
                    if (novaSenhaValue !== confirmaSenhaValue) {
                        if (confirmaSenhaErrorElement) { confirmaSenhaErrorElement.textContent = 'As senhas não coincidem.'; confirmaSenhaErrorElement.style.display = 'block'; } 
                        else { alert("As senhas não coincidem."); }
                        isFormValid = false;
                    }
                }
                if (isFormValid) {
                    if (typeof dadosPacienteAutenticado !== 'undefined' && dadosPacienteAutenticado && dadosPacienteAutenticado.online) {
                        dadosPacienteAutenticado.online.senha = novaSenhaValue;
                    }
                    window.location.href = 'redefinida-sucesso.html';
                } else {
                    if (novaSenhaErrorElement && novaSenhaErrorElement.style.display === 'block') { novaSenhaInputElement.focus(); } 
                    else if (confirmaSenhaErrorElement && confirmaSenhaErrorElement.style.display === 'block' && confirmaSenhaInputElement) { confirmaSenhaInputElement.focus(); } 
                    else { novaSenhaInputElement.focus(); }
                }
            });
        }
    })();

    // === Bloco para esqueci-senha.html ===
    (function() {
        if (document.getElementById('emailForm') && document.getElementById('emailInput')) {
            // Código específico para esqueci-senha.html (se houver)
        }
    })();

    // === Bloco genérico (script.js) === // (Este bloco estava vazio no seu original)
    (function() {
        // Código genérico (se houver)
    })();

    // === BLOCO PARA A PÁGINA DE LOGIN (ex: escolher-tipo-usuario.html) ===
    (function() {
        const formLoginUsuario = document.getElementById('formLoginUsuario');
        if (formLoginUsuario) { /* ... seu código ... */ 
            const mensagemErroLoginEl = document.getElementById('mensagem-login-erro'); // corrigido nome da var
            formLoginUsuario.addEventListener('submit', function(event) {
                event.preventDefault();
                const tipoSelecionadoRadio = document.querySelector('input[name="tipo"]:checked');
                if (!tipoSelecionadoRadio) {
                    if (mensagemErroLoginEl) { mensagemErroLoginEl.textContent = "Por favor, selecione o tipo de usuário."; mensagemErroLoginEl.style.display = 'block';} // corrigido nome da var
                    return;
                }
                const tipoSelecionado = tipoSelecionadoRadio.value;
                if (mensagemErroLoginEl) { mensagemErroLoginEl.textContent = ''; mensagemErroLoginEl.style.display = 'none'; }
                if (tipoSelecionado === 'beneficiario') {
                    const loginInput = document.getElementById('beneficiario-login');
                    const senhaInput = document.getElementById('beneficiario-senha');
                    if (!loginInput || !senhaInput) { console.error("Inputs beneficiário não encontrados"); return; }
                    const loginDigitado = loginInput.value.trim(); const senhaDigitada = senhaInput.value.trim();
                    if (loginDigitado === "" || senhaDigitada === "") {
                        if(mensagemErroLoginEl) { mensagemErroLoginEl.textContent = "Por favor, preencha login e senha."; mensagemErroLoginEl.style.display = 'block'; }
                        return;
                    }
                    if (dadosPacienteAutenticado && dadosPacienteAutenticado.online && loginDigitado === dadosPacienteAutenticado.online.login && senhaDigitada === dadosPacienteAutenticado.online.senha) {
                        localStorage.setItem('pacienteLogadoIdentificador', dadosPacienteAutenticado.online.login);
                        window.location.href = "INICIO.html";
                    } else {
                        if(mensagemErroLoginEl) { mensagemErroLoginEl.textContent = "Login ou senha incorretos para beneficiário."; mensagemErroLoginEl.style.display = 'block'; }
                    }
                } else if (tipoSelecionado === 'instituicao') {
                    const cnesInput = document.getElementById('instituicao-cnes');
                    const crmInput = document.getElementById('instituicao-crm');
                    const senhaInstInput = document.getElementById('instituicao-senha-inst');
                    if (!cnesInput || !crmInput || !senhaInstInput) { console.error("Inputs instituição não encontrados"); return; }
                    const cnesDigitado = cnesInput.value.trim(); const crmDigitado = crmInput.value.trim(); const senhaInstDigitada = senhaInstInput.value.trim();
                    if (cnesDigitado === "" || crmDigitado === "" || senhaInstDigitada === "") {
                        if(mensagemErroLoginEl) { mensagemErroLoginEl.textContent = "Por favor, preencha CNES, CRM/COREN e Senha."; mensagemErroLoginEl.style.display = 'block'; }
                        return;
                    }
                    if (dadosInstituicaoAutenticada && dadosInstituicaoAutenticada.online && cnesDigitado === dadosInstituicaoAutenticada.cnes && crmDigitado === dadosInstituicaoAutenticada.crmCorenResponsavel && senhaInstDigitada === dadosInstituicaoAutenticada.online.senha) {
                        localStorage.setItem('instituicaoLogadaIdentificador', dadosInstituicaoAutenticada.cnes);
                        window.location.href = "iniciomedico.html";
                    } else {
                        if(mensagemErroLoginEl) { mensagemErroLoginEl.textContent = "CNES, CRM/COREN ou Senha incorretos para instituição."; mensagemErroLoginEl.style.display = 'block'; }
                    }
                }
            });
        }
    })();

    // === BLOCO PARA A PÁGINA DE INÍCIO/DASHBOARD DO PACIENTE (ex: INICIO.html) ===
    (function() {
        if (document.body.id === 'paginaPerfilPaciente') { // NOTA: Este ID pode ter mudado para 'paginaUnificadaPaciente'
             /* ... seu código ... */ 
            const identificadorLogado = localStorage.getItem('pacienteLogadoIdentificador');
            if (dadosPacienteAutenticado && dadosPacienteAutenticado.online && identificadorLogado && identificadorLogado === dadosPacienteAutenticado.online.login) {
                const elSaudacaoNome = document.getElementById('saudacaoNome');
                if (elSaudacaoNome) {
                    elSaudacaoNome.textContent = dadosPacienteAutenticado.primeiroNome || (dadosPacienteAutenticado.nomeCompleto ? dadosPacienteAutenticado.nomeCompleto.split(' ')[0] : 'Usuário');
                }
                const btnPainelMeuPerfil = document.getElementById('btnPainelMeuPerfil');
                if (btnPainelMeuPerfil) {
                    btnPainelMeuPerfil.addEventListener('click', function() { window.location.href = 'alterar-perfil.html'; });
                }
            }
        }
    })();

    // === BLOCO PARA A PÁGINA DE CADASTRO DE PACIENTE (ex: iniciomedico.html) ===
    (function() {
        if (document.body.id === 'paginaCadastroPaciente') { // NOTA: Verifique se este ID é de uma página que ainda existe ou se a lógica foi movida
            /* ... seu código de máscaras e validação do formulário de cadastro de paciente ... */ 
            const nomeMedicoEl = document.getElementById('nomeMedicoLogado');
            const instituicaoLogadaCNES = localStorage.getItem('instituicaoLogadaIdentificador');
            if (nomeMedicoEl && instituicaoLogadaCNES && typeof dadosInstituicaoAutenticada !== 'undefined' && dadosInstituicaoAutenticada.cnes === instituicaoLogadaCNES) {
                nomeMedicoEl.textContent = dadosInstituicaoAutenticada.nomeInstituicao || "Doutor(a)";
            } else if (nomeMedicoEl) {
                nomeMedicoEl.textContent = "Joana"; // Fallback
            }

            function formatCPF(cpf) { cpf = cpf.replace(/\D/g, "").slice(0, 11); cpf = cpf.replace(/(\d{3})(\d)/, "$1.$2"); cpf = cpf.replace(/(\d{3})(\d)/, "$1.$2"); cpf = cpf.replace(/(\d{3})(\d{1,2})$/, "$1-$2"); return cpf; }
            function formatDate(date) { date = date.replace(/\D/g, "").slice(0, 8); date = date.replace(/(\d{2})(\d)/, "$1/$2"); date = date.replace(/(\d{2})(\d)/, "$1/$2"); return date; }
            function formatCEP(cep) { cep = cep.replace(/\D/g, "").slice(0, 8); cep = cep.replace(/^(\d{5})(\d)/, "$1-$2"); return cep; }
            function formatPhone(phone) { phone = phone.replace(/\D/g, "").slice(0, 11); if (phone.length > 10) { phone = phone.replace(/^(\d{2})(\d{5})(\d{4}).*/, "($1) $2-$3"); } else if (phone.length > 9) { phone = phone.replace(/^(\d{2})(\d{4})(\d{4}).*/, "($1) $2-$3"); } else if (phone.length > 5) { phone = phone.replace(/^(\d{2})(\d{0,4})/, "($1) $2"); } else if (phone.length === 1 && phone !== "("){ phone = "("+phone; } return phone; }
            function formatCartaoSUS(sus) { sus = sus.replace(/\D/g, "").slice(0, 15); sus = sus.replace(/(\d{3})(\d)/, "$1 $2"); sus = sus.replace(/(\d{4})(\d)/, "$1 $2"); sus = sus.replace(/(\d{4})(\d)/, "$1 $2"); return sus; }

            const camposParaMascarar = [
                { id: 'pacienteCPF', func: formatCPF, maxlength: 14 }, { id: 'pacienteDataNascimento', func: formatDate, maxlength: 10 },
                { id: 'pacienteCEP', func: formatCEP, maxlength: 9 }, { id: 'pacienteCelular', func: formatPhone, maxlength: 15 },
                { id: 'pacienteTelefone', func: formatPhone, maxlength: 14 }, { id: 'pacienteCartaoSUS', func: formatCartaoSUS, maxlength: 19 }
            ];
            camposParaMascarar.forEach(campoInfo => {
                const inputElement = document.getElementById(campoInfo.id);
                if (inputElement) {
                    if(campoInfo.maxlength) inputElement.maxLength = campoInfo.maxlength;
                    inputElement.addEventListener('input', function (e) { e.target.value = campoInfo.func(e.target.value); });
                }
            });

            const racaSelect = document.getElementById('pacienteRacaSelect');
            const racaOutrosInput = document.getElementById('pacienteRacaOutros');
            const containerRacaOutros = document.getElementById('containerRacaOutros');
            if (racaSelect && racaOutrosInput && containerRacaOutros) {
                racaSelect.addEventListener('change', function() {
                    if (this.value === 'Outros') { containerRacaOutros.style.display = 'block'; racaOutrosInput.focus(); } 
                    else { containerRacaOutros.style.display = 'none'; racaOutrosInput.value = ''; }
                });
            }
            
            const formCadastro = document.getElementById('formCadastroPaciente');
            if (formCadastro) {
                const camposObrigatorios = [
                    { inputId: 'pacienteNome', erroId: 'erro-pacienteNome', nomeCampo: 'Nome Completo' }, { inputId: 'pacienteCPF', erroId: 'erro-pacienteCPF', nomeCampo: 'CPF' },
                    { inputId: 'pacienteDataNascimento', erroId: 'erro-pacienteDataNascimento', nomeCampo: 'Data de Nascimento' }, { inputId: 'pacienteCartaoSUS', erroId: 'erro-pacienteCartaoSUS', nomeCampo: 'Cartão SUS' },
                    { inputId: 'pacienteCEP', erroId: 'erro-pacienteCEP', nomeCampo: 'CEP' }, { inputId: 'pacienteCelular', erroId: 'erro-pacienteCelular', nomeCampo: 'Celular (com DDD)' },
                    { inputId: 'pacienteNacionalidade', erroId: 'erro-pacienteNacionalidade', nomeCampo: 'Nacionalidade' }, { inputId: 'pacienteEstadoSelect', erroId: 'erro-pacienteEstadoSelect', nomeCampo: 'Estado (UF)' }
                ];
                function toggleErroCampo(campoId, erroId, mostrar, mensagem = "Este dado é obrigatório.") {
                    const campoEl = document.getElementById(campoId); const erroEl = document.getElementById(erroId);
                    if (erroEl) { erroEl.textContent = mensagem; erroEl.style.display = mostrar ? 'block' : 'none'; }
                    if (campoEl) { if (mostrar) campoEl.classList.add('campo-invalido'); else campoEl.classList.remove('campo-invalido'); }
                }
                camposObrigatorios.forEach(config => {
                    const elemento = document.getElementById(config.inputId);
                    if (elemento) {
                        const evento = elemento.tagName.toLowerCase() === 'select' ? 'change' : 'input';
                        elemento.addEventListener(evento, () => { toggleErroCampo(config.inputId, config.erroId, false); });
                    }
                });
                if (racaOutrosInput) { racaOutrosInput.addEventListener('input', () => { toggleErroCampo('pacienteRacaOutros', 'erro-pacienteRacaOutros', false); }); }

                formCadastro.addEventListener('submit', function(event) {
                    event.preventDefault(); let todosCamposValidos = true;
                    camposObrigatorios.forEach(config => { toggleErroCampo(config.inputId, config.erroId, false); });
                    if(document.getElementById('erro-pacienteRacaOutros')) toggleErroCampo('pacienteRacaOutros', 'erro-pacienteRacaOutros', false);
                    camposObrigatorios.forEach(config => {
                        const elemento = document.getElementById(config.inputId);
                        if (elemento && elemento.value.trim() === '') { toggleErroCampo(config.inputId, config.erroId, true, `Campo obrigatório.`); todosCamposValidos = false; }
                    });
                    if (racaSelect && racaSelect.value === 'Outros') {
                        if (racaOutrosInput && racaOutrosInput.value.trim() === '') { toggleErroCampo('pacienteRacaOutros', 'erro-pacienteRacaOutros', true, 'Por favor, especifique a raça.'); todosCamposValidos = false; }
                    }
                    if (!todosCamposValidos) {
                        alert("Por favor, preencha todos os campos obrigatórios destacados.");
                        const primeiroInvalido = document.querySelector('.campo-invalido'); if(primeiroInvalido) primeiroInvalido.focus(); return;
                    }
                    const confirmacao = window.confirm("Os dados do paciente estão corretos e você deseja prosseguir com o cadastro?");
                    if (confirmacao) {
    alert("Paciente cadastrado com sucesso!");
    formCadastro.reset();
    if (containerRacaOutros) containerRacaOutros.style.display = 'none';
    if (racaOutrosInput) racaOutrosInput.value = '';
    camposObrigatorios.forEach(config => {
        toggleErroCampo(config.inputId, config.erroId, false);
    });
    if (document.getElementById('erro-pacienteRacaOutros')) {
        toggleErroCampo('pacienteRacaOutros', 'erro-pacienteRacaOutros', false);
    }
    const primeiroCampoFocavel = formCadastro.querySelector('input:not([type=hidden]), select, textarea');
    if (primeiroCampoFocavel) primeiroCampoFocavel.focus();

    // Redirecionar após o sucesso
    window.location.href = "iniciomedico.html";
}

                });
            }
        }
    })();

    console.log("geral.js: IIFE (Função Auto-Executável) finalizada.");
})(); // Fim da IIFE Principal
