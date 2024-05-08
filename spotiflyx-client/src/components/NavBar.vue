<template>
    <nav>
        <a class="logo">
            <RouterLink to="/home"><img src="/logo_trans.png" alt="logo"></RouterLink>
        </a>
        <div class="btns">
            <RouterLink to="/home"><button>Home</button></RouterLink>
            <RouterLink to="/favorites"><button>Favoris</button></RouterLink>
            <RouterLink to="/account"><button>Mon compte</button></RouterLink>
        </div>
    </nav>
    <div class="shadow"></div>
</template>

<style scoped>
nav {
    width: 100%;
    height: 8vh;
    background-color: transparent;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.shadow {
    width: 100%;
    height: 0px;
    box-shadow: 0px -2px 20px 6px var(--primary-color);
}

.logo {
    width: 10vw;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}

.logo a {
    text-decoration: none;
    display: flex;
    justify-content: center;
    align-items: center;
}

.logo img {
    width: auto;
    height: 6vh;
}

.btns {
    width: 50vw;
    height: 100%;
    display: flex;
    justify-content: space-evenly;
    align-items: center;
}

button {
    width: 10vw;
    height: 60%;
    background-color: transparent;
    border: none;
    font-size: 1rem;
    font-weight: 700;
    color: var(--text-color);
    cursor: pointer;
}
</style>