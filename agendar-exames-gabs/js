document.addEventListener('DOMContentLoaded', () => {
    const btnAgendarExames = document.getElementById('btnAgendarExames');
    const dateModal = document.getElementById('dateModal');
    const examModal = document.getElementById('examModal');

    const currentDayDisplay = document.querySelector('.current-date-display');
    const prevMonthBtn = document.querySelector('.prev-month');
    const nextMonthBtn = document.querySelector('.next-month');
    const monthYearDisplay = document.querySelector('.month-year');
    const calendarGrid = document.querySelector('.calendar-grid');
    const cancelDateSelectionBtn = document.getElementById('cancelDateSelection');
    const okDateSelectionBtn = document.getElementById('okDateSelection');

    let currentDate = new Date();
    let activeDate = new Date(); // Dia inicialmente selecionado será o dia atual
    let currentMonth = currentDate.getMonth();
    let currentYear = currentDate.getFullYear();

    const monthNames = [
        "Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
        "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"
    ];
    const dayNames = [
        "Domingo", "Segunda-feira", "Terça-feira", "Quarta-feira",
        "Quinta-feira", "Sexta-feira", "Sábado"
    ];

    function openModal(modal) {
        modal.classList.add('active');
    }

    function closeModal(modal) {
        modal.classList.remove('active');
    }

    function renderCalendar() {
        while (calendarGrid.children.length > 7) {
            calendarGrid.removeChild(calendarGrid.lastChild);
        }

        monthYearDisplay.textContent = `${monthNames[currentMonth]} ${currentYear}`;

        const firstDayOfMonth = new Date(currentYear, currentMonth, 1).getDay();
        const lastDayOfMonth = new Date(currentYear, currentMonth + 1, 0).getDate();

        for (let i = 0; i < firstDayOfMonth; i++) {
            const emptyDay = document.createElement('span');
            emptyDay.classList.add('day', 'empty');
            calendarGrid.appendChild(emptyDay);
        }

        for (let day = 1; day <= lastDayOfMonth; day++) {
            const dayElement = document.createElement('span');
            dayElement.classList.add('day');
            dayElement.textContent = day;

            if (day === currentDate.getDate() &&
                currentMonth === currentDate.getMonth() &&
                currentYear === currentDate.getFullYear()) {
                dayElement.classList.add('today');
            }

            if (day === activeDate.getDate() &&
                currentMonth === activeDate.getMonth() &&
                currentYear === activeDate.getFullYear()) {
                dayElement.classList.add('selected');
            }

            dayElement.addEventListener('click', () => {
                const previouslySelected = calendarGrid.querySelector('.day.selected');
                if (previouslySelected) {
                    previouslySelected.classList.remove('selected');
                }
                dayElement.classList.add('selected');

                activeDate = new Date(currentYear, currentMonth, day);
                updateCurrentDateDisplay(activeDate);
            });

            calendarGrid.appendChild(dayElement);
        }
    }

    function updateCurrentDateDisplay(date) {
        const dayOfWeek = dayNames[date.getDay()];
        const dayOfMonth = date.getDate();
        const month = monthNames[date.getMonth()];
        const year = date.getFullYear();
        currentDayDisplay.textContent = `${dayOfWeek}, ${dayOfMonth} de ${month} de ${year}`;
    }

    if (btnAgendarExames) {
        btnAgendarExames.addEventListener('click', () => {
            renderCalendar();
            updateCurrentDateDisplay(activeDate);
            openModal(dateModal);
        });
    }

    if (prevMonthBtn) {
        prevMonthBtn.addEventListener('click', () => {
            currentMonth--;
            if (currentMonth < 0) {
                currentMonth = 11;
                currentYear--;
            }
            renderCalendar();
        });
    }

    if (nextMonthBtn) {
        nextMonthBtn.addEventListener('click', () => {
            currentMonth++;
            if (currentMonth > 11) {
                currentMonth = 0;
                currentYear++;
            }
            renderCalendar();
        });
    }

    if (cancelDateSelectionBtn) {
        cancelDateSelectionBtn.addEventListener('click', () => {
            closeModal(dateModal);
        });
    }

    if (okDateSelectionBtn) {
        okDateSelectionBtn.addEventListener('click', () => {
            closeModal(dateModal);
            openModal(examModal);
        });
    }

    if (dateModal) {
        dateModal.addEventListener('click', (e) => {
            if (e.target === dateModal) {
                closeModal(dateModal);
            }
        });
    }

    if (examModal) {
        examModal.addEventListener('click', (e) => {
            if (e.target === examModal) {
                closeModal(examModal);
            }
        });

        const radioOptions = examModal.querySelectorAll('input[name="examType"]');
        radioOptions.forEach(radio => {
            radio.addEventListener('change', () => {
                // console.log(`Exame selecionado: ${radio.value}`);
            });
        });
    }

    renderCalendar();
    updateCurrentDateDisplay(activeDate);
});
