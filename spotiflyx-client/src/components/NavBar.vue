<template>
    <nav>
        <a class="logo">
            <RouterLink to="/"><img src="/logo_trans.png" alt="logo"></RouterLink>
        </a>
        <div class="btns">
            <button>Home</button>
            <button>Favorites</button>
            <button>Account</button>
        </div>
    </nav>
</template>

<style scoped>
nav {
    width: 100%;
    height: 8vh;
    background-color: var(--vt-c-indigo);
    display: flex;
    justify-content: space-between;
    align-items: center;
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