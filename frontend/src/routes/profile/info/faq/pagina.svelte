<script>
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
                Om muziek te delen vanuit Deezer, volg deze stappen:
                <ol>
                    <li>Open de Deezer-app op je mobiele apparaat of browser.</li>
                    <li>Zoek het nummer dat je wilt delen.</li>
                    <li>Klik op de deelknop naast het nummer.</li>
                    <li>Kies "Harmonize" als optie om het nummer te delen.</li>
                </ol>
            </div>
        {/if}
    </div>

    <!-- Voeg meer FAQ-items toe op dezelfde manier voor andere vragen -->
</Panel>
