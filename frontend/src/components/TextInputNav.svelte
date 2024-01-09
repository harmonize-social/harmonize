<script lang="ts">
    import { get, throwError } from '../fetch';

    export let placeholder = 'Search';
    let input: string = '';
    let list: string[] = [];

    function handleTextInput(event: KeyboardEvent) {
        if (event.key === 'Enter') {
            handleInput();
        }
    }

    async function handleInput() {
        try {
            const response = (await get('/search?username=' + input)) as any;
            if (response.error) throwError(response.error);
            list = response;
        } catch (e) {
            console.log(e);
            throwError('Internal server error');
        }
        console.log(list);
    }
</script>

<div class="search-container">
    <input
        type="text"
        class="textInputNav"
        {placeholder}
        on:keydown={handleTextInput}
        bind:value={input}
    />
    {#if list.length > 0}
        <div class="list">
            {#each list as item}
                <a href="/profile/{item}">{item}</a>
            {/each}
        </div>
    {/if}
</div>

<style>
    .search-container {
        position: relative;
        display: inline-block;
    }

    .textInputNav {
        grid-area: searchbox;


        width: 200px;
        height: 55px;
        border-radius: 50px;
        border: 1px solid black;
        color: rebeccapurple;
    }

    .list {
        display: flex;
        align-items: center;
        flex-direction: column;
        position: absolute;
        top: 100%;
        left: 0;
        width: 200px;
        border: 1px solid #ddd;
        border-radius: 4px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
        background-color: white;
        z-index: 1000;
    }

    .list a {
        cursor: pointer;
        text-decoration: none;
    }
    .list a:hover {
        text-transform: uppercase;
    }
</style>
