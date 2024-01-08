<script>
    export let title = 'Panel';
</script>

<div class="panel">
    <h2>{title}</h2>
    <slot />
</div>

<style>
    .panel {
        border-radius: 10px;
        padding: 20px;
        box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
        display: flex;
        flex-direction: column;
        flex-wrap: wrap;
        align-items: center;
    }

    .panel h2 {
        flex: 0 0 auto;
    }

    .panel slot {
        flex: 1 1 auto;
        overflow: auto;
    }
</style>
