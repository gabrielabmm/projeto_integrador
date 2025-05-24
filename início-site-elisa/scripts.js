// scripts.js

document.addEventListener('DOMContentLoaded', function() {
    const cookieBanner = document.getElementById('cookie-banner');
    const acceptButton = document.getElementById('cookie-accept');
    const declineButton = document.getElementById('cookie-decline');

    // Função para definir um cookie (ou localStorage)
    function setCookiePreference(value, days) {
        // Usando localStorage para simplicidade. Para cookies reais com expiração,
        // a lógica é um pouco mais complexa.
        localStorage.setItem('cookieConsent', value);
        // Se quisesse usar cookies HTTP:
        // let expires = "";
        // if (days) {
        //     const date = new Date();
        //     date.setTime(date.getTime() + (days*24*60*60*1000));
        //     expires = "; expires=" + date.toUTCString();
        // }
        // document.cookie = "cookieConsent=" + (value || "")  + expires + "; path=/";
    }

    // Função para obter a preferência de cookie
    function getCookiePreference() {
        return localStorage.getItem('cookieConsent');
        // Se usasse cookies HTTP:
        // const nameEQ = "cookieConsent=";
        // const ca = document.cookie.split(';');
        // for(let i=0;i < ca.length;i++) {
        //     let c = ca[i];
        //     while (c.charAt(0)==' ') c = c.substring(1,c.length);
        //     if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
        // }
        // return null;
    }

    // Verificar se o consentimento já foi dado
    if (cookieBanner) {
        cookieBanner.style.display = 'block'; // Ou 'flex' se você usou display:flex para o layout interno do banner
                                              // Verifique qual display é usado na regra .cookie-banner no seu CSS
                                              // quando ele deve estar visível.
    }

    // 2. Modifique os botões para APENAS esconder o banner, sem salvar a preferência.
    if (acceptButton && cookieBanner) {
        acceptButton.addEventListener('click', function() {
            cookieBanner.style.display = 'none'; // Apenas esconde o banner
            console.log("Cookies 'aceitos' para esta visualização da página.");
            // Nenhuma chamada para localStorage.setItem() aqui
        });
    }

    if (declineButton && cookieBanner) {
        declineButton.addEventListener('click', function() {
            cookieBanner.style.display = 'none'; // Apenas esconde o banner
            console.log("Cookies 'recusados' para esta visualização da página.");
            // Nenhuma chamada para localStorage.setItem() aqui
        });
    }
});
