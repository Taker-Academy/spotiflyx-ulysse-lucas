<template>
    <footer>
        <div class="footer">
            <div class="footer__left">
                <div class="footer__left__logo">
                    <RouterLink to="/home"><img src="/logo_trans.png" alt="logo"></RouterLink>
                </div>
                <div class="footer__left__text">
                    <p>Spotiflyx</p>
                    <p>© 2024 Spotiflyx</p>
                </div>
            </div>
            <div class="footer__right">
                <div class="footer__right__text">
                    <RouterLink to="/home"><p>Home</p></RouterLink>
                    <RouterLink to="/account"><p>Mon compte</p></RouterLink>
                    <RouterLink to="/favorites"><p>Mes favoris</p></RouterLink>
                </div>
            </div>
        </div>
    </footer>
</template>

<style scoped>
footer {
    width: 100%;
    height: 12vh;
    background-color: var(--primary-color);
    display: flex;
    justify-content: center;
    align-items: center;
}

.footer {
    width: 80%;
    height: 90%;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.footer__left {
    display: flex;
    align-items: center;
}

.footer__left__logo {
    width: 50px;
    height: 50px;
    margin-right: 1rem;
}

.footer__left__logo img {
    width: 100%;
    height: 100%;
}

.footer__left__text {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.footer__right {
    display: flex;
    align-items: center;
}

.footer__right__text {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.footer__right__text p {
    margin-bottom: 0.5rem;
    color: var(--color-background-soft);
    font-weight: 700;
    text-decoration: none;
    font-style: normal;
    text-decoration-line: none;
}

.footer__left__text p {
    margin-bottom: 0.5rem;
    color: var(--color-background-soft);
    font-weight: 700;
}

</style>