let map;
let activeMarkers = [];
const MAPTILER_API_KEY = 'eR03qKIxxAJFEfbzwqF5'; // Sua chave MapTiler

// Função para inicializar o mapa Leaflet
function initializeLeafletMap() {
    const initialCoords = [-14.235004, -51.92528]; // Centro do Brasil
    map = L.map('map').setView(initialCoords, 4);

    L.tileLayer(`https://api.maptiler.com/maps/streets-v2/{z}/{x}/{y}.png?key=${MAPTILER_API_KEY}`, {
        attribution: '<a href="https://www.maptiler.com/copyright/" target="_blank">&copy; MapTiler</a> <a href="https://www.openstreetmap.org/copyright" target="_blank">&copy; OpenStreetMap contributors</a>',
        maxZoom: 19,
    }).addTo(map);
    console.log("Mapa base MapTiler inicializado.");
}

// Função para limpar marcadores antigos do mapa
function clearMarkers() {
    activeMarkers.forEach(marker => map.removeLayer(marker));
    activeMarkers = [];
    const ubsList = document.getElementById('ubs-list');
    if (ubsList) ubsList.innerHTML = ''; // Limpa a lista também, caso houvesse algo
}

// Função para criar marcadores Leaflet no mapa
function createLeafletMarker(lat, lon, name, vicinity) {
    const marker = L.marker([lat, lon]).addTo(map)
        .bindPopup(`<strong>${name}</strong><br>${vicinity || ''}`);
    activeMarkers.push(marker);
    marker.openPopup(); // Abre o popup para o local do CEP
}

// Função para buscar endereço por CEP (usando API ViaCEP)
async function buscarEnderecoPorCEP(cep) {
    const cepFormatado = cep.replace(/\D/g, '');
    if (cepFormatado.length !== 8) {
        alert("CEP inválido. Por favor, digite 8 números.");
        return null;
    }
    try {
        const response = await fetch(`https://viacep.com.br/ws/${cepFormatado}/json/`);
        if (!response.ok) throw new Error('Não foi possível buscar o CEP via ViaCEP.');
        const data = await response.json();
        if (data.erro) {
            alert("CEP não encontrado pelo ViaCEP.");
            return null;
        }
        return `${data.logradouro || ''}, ${data.bairro || ''}, ${data.localidade || ''} - ${data.uf || ''}`;
    } catch (error) {
        console.error("Erro ao buscar CEP no ViaCEP:", error);
        alert(error.message);
        return null;
    }
}

// Função para geocodificar o endereço (com MapTiler) e centralizar o mapa
async function geocodeAddressAndCenterMap(addressString, cepInformado) {
    if (!map) {
        alert("O mapa ainda não foi carregado.");
        return;
    }
    if (!MAPTILER_API_KEY || MAPTILER_API_KEY === 'SUA_CHAVE_API_MAPTILER') { // Verificação genérica
        alert("Chave de API do MapTiler não configurada ou inválida para esta demonstração.");
        // Mesmo com a chave fornecida, o serviço de geocodificação pode falhar se a chave não tiver permissão.
        // Mas vamos tentar carregar os tiles do mapa pelo menos.
    }
    clearMarkers();

    console.log(`Tentando geocodificar endereço: ${addressString} com a chave ${MAPTILER_API_KEY}`);

    try {
        // 1. Geocodificar o endereço para obter coordenadas
        // Esta chamada PODE falhar com "Invalid key" se a chave 'eR03qKIxxAJFEfbzwqF5'
        // não tiver permissão nem mesmo para o serviço básico de geocodificação de endereços.
        const geocodeUrl = `https://api.maptiler.com/geocoding/${encodeURIComponent(addressString)}.json?key=${MAPTILER_API_KEY}&language=pt&limit=1`;
        
        const geoResponse = await fetch(geocodeUrl);
        
        if (!geoResponse.ok) {
            // Se falhar aqui, é provável que a chave seja inválida também para este serviço.
            const errorStatus = geoResponse.status;
            const errorData = await geoResponse.json().catch(() => null); // Tenta pegar o corpo do erro
            console.error(`Falha ao geocodificar endereço com MapTiler. Status: ${errorStatus}`, errorData);
            alert(`Não foi possível obter as coordenadas para o endereço. \nO MapTiler retornou um erro ${errorStatus}. Verifique o console para detalhes.\nIsso pode ser devido a uma chave de API inválida para o serviço de geocodificação.`);
            // Centraliza no Brasil como fallback se a geocodificação falhar
            map.setView([-14.235004, -51.92528], 4);
            return;
        }
        
        const geoData = await geoResponse.json();

        if (geoData.features && geoData.features.length > 0) {
            const coordinates = geoData.features[0].center; // Formato [longitude, latitude]
            const lat = coordinates[1];
            const lon = coordinates[0];

            map.setView([lat, lon], 15); // Centraliza o mapa e aplica zoom
            createLeafletMarker(lat, lon, `Local do CEP: ${cepInformado}`, addressString);
            console.log(`Endereço geocodificado e mapa centralizado em: ${lat}, ${lon}`);
        } else {
            alert('Não foi possível encontrar a localização exata para o endereço/CEP informado via MapTiler.');
            map.setView([-14.235004, -51.92528], 4); // Fallback
        }
    } catch (error) {
        console.error("Erro no processo de geocodificação do endereço:", error);
        // Este catch pode pegar erros de rede ou outros problemas antes da verificação do geoResponse.ok
        alert(`Ocorreu um erro ao tentar geocodificar o endereço: ${error.message}`);
        map.setView([-14.235004, -51.92528], 4); // Fallback
    }

    // AVISO: A busca por UBS foi removida desta função, pois
    // a chave de API estava apresentando erro "Invalid key" para esse tipo de serviço.
    // Esta versão foca em exibir o mapa e centralizar no CEP.
    const ubsList = document.getElementById('ubs-list');
    if (ubsList) {
        ubsList.innerHTML = '<li>A busca por UBS próximas está desabilitada no momento devido a limitações da chave de API para este serviço.</li>';
    }
}

// Adiciona listeners quando o DOM estiver carregado
document.addEventListener('DOMContentLoaded', () => {
    initializeLeafletMap(); 

    const cepForm = document.getElementById('cep-form');
    const cepInput = document.getElementById('cep');

    if (cepForm && cepInput) {
        cepForm.addEventListener('submit', async (event) => {
            event.preventDefault();
            const cepValue = cepInput.value;
            if (!cepValue.trim()) {
                alert("Por favor, insira um CEP.");
                return;
            }
            const endereco = await buscarEnderecoPorCEP(cepValue);
            if (endereco) {
                if (map) {
                    await geocodeAddressAndCenterMap(endereco, cepValue);
                } else {
                    alert("O mapa não está carregado. Tente recarregar a página.");
                }
            }
        });

        cepInput.addEventListener('input', (e) => {
            let value = e.target.value.replace(/\D/g, '');
            if (value.length > 5) {
                value = value.substring(0, 5) + '-' + value.substring(5, 8);
            }
            e.target.value = value.substring(0, 9);
        });
    } else {
        console.error("Formulário de CEP ou campo de input não encontrados.");
    }
});
