const cards = document.querySelectorAll('.card');

if (cards.length > 4) {
    for (let i = 4; i < cards.length; i++) {
        cards[i].style.display = 'none';
    }
}