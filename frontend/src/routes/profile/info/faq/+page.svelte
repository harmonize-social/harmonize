<script>
	import { onMount } from 'svelte';

// @ts-nocheck

    import NavBar from '../../../../components/NavBar.svelte';
    import Panel from '../../../../components/Panel.svelte';

    let visible = {};

    function toggleAnswer(id) {
        visible[id] = !visible[id];
    }

</script>

<style>
    /* Stijlen voor de FAQ-pagina */
    .faq-item {
        margin-bottom: 20px;
    }
   
    .answer {
        display: none;
        margin-bottom: 10px;
    }
    .visible {
        display: block;
    }
</style>

<NavBar current_page="/info/faq"></NavBar>
<Panel title="Frequently Asked Questions">
    <div class="faq-item">
        {#if visible['q1']}
            <div class="answer visible" id="q1">
                To share music via Deezer, follow these steps:
                <ol>
                    <li>Go to your profile page</li>
                    <li>Click on the library tab</li>
                    <li>Click on the musical item you want to share</li>
                    <li>Choose a caption and share!</li>
                </ol>
            </div>
        {/if}
    </div>

    <!-- Voeg meer FAQ-items toe op dezelfde manier voor andere vragen -->
</Panel>
