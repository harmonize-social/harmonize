<script lang="ts">
    import { onMount } from 'svelte';

    export let previewUrl: string = '';
    let playerButton: HTMLElement;
    let audio: HTMLAudioElement;
    let timeline: HTMLElement;
    let soundButton: HTMLElement;

    const playIcon = `
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="#3D3132">
    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd" />
  </svg>
      `;

    const pauseIcon = `
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="#3D3132">
  <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zM7 8a1 1 0 012 0v4a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v4a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" />
</svg>
      `;

    const soundIcon = `
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="#3D3132">
  <path fill-rule="evenodd" d="M9.383 3.076A1 1 0 0110 4v12a1 1 0 01-1.707.707L4.586 13H2a1 1 0 01-1-1V8a1 1 0 011-1h2.586l3.707-3.707a1 1 0 011.09-.217zM14.657 2.929a1 1 0 011.414 0A9.972 9.972 0 0119 10a9.972 9.972 0 01-2.929 7.071 1 1 0 01-1.414-1.414A7.971 7.971 0 0017 10c0-2.21-.894-4.208-2.343-5.657a1 1 0 010-1.414zm-2.829 2.828a1 1 0 011.415 0A5.983 5.983 0 0115 10a5.984 5.984 0 01-1.757 4.243 1 1 0 01-1.415-1.415A3.984 3.984 0 0013 10a3.983 3.983 0 00-1.172-2.828 1 1 0 010-1.415z" clip-rule="evenodd" />
</svg>`;

    const muteIcon = `
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="#3D3132">
  <path fill-rule="evenodd" d="M9.383 3.076A1 1 0 0110 4v12a1 1 0 01-1.707.707L4.586 13H2a1 1 0 01-1-1V8a1 1 0 011-1h2.586l3.707-3.707a1 1 0 011.09-.217zM12.293 7.293a1 1 0 011.414 0L15 8.586l1.293-1.293a1 1 0 111.414 1.414L16.414 10l1.293 1.293a1 1 0 01-1.414 1.414L15 11.414l-1.293 1.293a1 1 0 01-1.414-1.414L13.586 10l-1.293-1.293a1 1 0 010-1.414z" clip-rule="evenodd" />
</svg>`;

    function toggleAudio() {
        if (audio.paused) {
            audio.play();
            playerButton.innerHTML = pauseIcon;
        } else {
            audio.pause();
            playerButton.innerHTML = playIcon;
        }
    }

    function changeTimelinePosition() {
        const percentagePosition = (100 * audio.currentTime) / audio.duration;
        timeline.style.backgroundSize = `${percentagePosition}% 100%`;
        timeline.value = percentagePosition;
    }

    onMount(() => {
        audio.ontimeupdate = changeTimelinePosition;
        audio.onended = audioEnded;
    });

    function audioEnded() {
        playerButton.innerHTML = playIcon;
    }

    function changeSeek() {
        const time = (timeline.value * audio.duration) / 100;
        audio.currentTime = time;
    }

    function toggleSound() {
        audio.muted = !audio.muted;
        soundButton.innerHTML = audio.muted ? muteIcon : soundIcon;
    }
</script>

<div class="audio-player">
    <audio src={previewUrl} bind:this={audio}></audio>
    <div class="controls">
        <button class="player-button" on:click={toggleAudio} bind:this={playerButton}>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="#3D3132">
                <path
                    fill-rule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z"
                    clip-rule="evenodd"
                />
            </svg>
        </button>
        <input
            type="range"
            class="timeline"
            max="100"
            value="0"
            on:change={changeSeek}
            bind:this={timeline}
        />
        <button class="sound-button" on:click={toggleSound} bind:this={soundButton}>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="#3D3132">
                <path
                    fill-rule="evenodd"
                    d="M9.383 3.076A1 1 0 0110 4v12a1 1 0 01-1.707.707L4.586 13H2a1 1 0 01-1-1V8a1 1 0 011-1h2.586l3.707-3.707a1 1 0 011.09-.217zM14.657 2.929a1 1 0 011.414 0A9.972 9.972 0 0119 10a9.972 9.972 0 01-2.929 7.071 1 1 0 01-1.414-1.414A7.971 7.971 0 0017 10c0-2.21-.894-4.208-2.343-5.657a1 1 0 010-1.414zm-2.829 2.828a1 1 0 011.415 0A5.983 5.983 0 0115 10a5.984 5.984 0 01-1.757 4.243 1 1 0 01-1.415-1.415A3.984 3.984 0 0013 10a3.983 3.983 0 00-1.172-2.828 1 1 0 010-1.415z"
                    clip-rule="evenodd"
                />
            </svg>
        </button>
    </div>
</div>

<style>
    .audio-player {
        --player-button-width: 3em;
        --sound-button-width: 2em;
        --space: 0.5em;
        width: 15rem;
    }

    .controls {
        display: flex;
        flex-direction: row;
        align-items: center;
        width: 100%;
        margin-top: 10px;
    }

    .player-button {
        background-color: transparent;
        border: 0;
        width: var(--player-button-width);
        height: var(--player-button-width);
        cursor: pointer;
        padding: 0;
    }

    .timeline {
        -webkit-appearance: none;
        width: calc(100% - (var(--player-button-width) + var(--sound-button-width) + var(--space)));
        height: 0.5em;
        background-color: #e5e5e5;
        border-radius: 5px;
        background-size: 0% 100%;
        background-image: linear-gradient(#de5e97, #de5e97);
        background-repeat: no-repeat;
        margin-right: var(--space);
    }

    .timeline::-webkit-slider-thumb {
        -webkit-appearance: none;
        width: 1em;
        height: 1em;
        border-radius: 50%;
        cursor: pointer;
        opacity: 0;
        transition: all 0.1s;
        background-color: #a94672;
    }

    .timeline::-moz-range-thumb {
        -webkit-appearance: none;
        width: 1em;
        height: 1em;
        border-radius: 50%;
        cursor: pointer;
        opacity: 0;
        transition: all 0.1s;
        background-color: #a94672;
    }

    .timeline::-ms-thumb {
        -webkit-appearance: none;
        width: 1em;
        height: 1em;
        border-radius: 50%;
        cursor: pointer;
        opacity: 0;
        transition: all 0.1s;
        background-color: #a94672;
    }

    .timeline::-webkit-slider-thumb:hover {
        background-color: #943f65;
    }

    .timeline:hover::-webkit-slider-thumb {
        opacity: 1;
    }

    .timeline::-moz-range-thumb:hover {
        background-color: #943f65;
    }

    .timeline:hover::-moz-range-thumb {
        opacity: 1;
    }

    .timeline::-ms-thumb:hover {
        background-color: #943f65;
    }

    .timeline:hover::-ms-thumb {
        opacity: 1;
    }

    .timeline::-webkit-slider-runnable-track {
        -webkit-appearance: none;
        box-shadow: none;
        border: none;
        background: transparent;
    }

    .timeline::-moz-range-track {
        -webkit-appearance: none;
        box-shadow: none;
        border: none;
        background: transparent;
    }

    .timeline::-ms-track {
        -webkit-appearance: none;
        box-shadow: none;
        border: none;
        background: transparent;
    }

    .sound-button {
        background-color: transparent;
        border: 0;
        width: var(--sound-button-width);
        height: var(--sound-button-width);
        cursor: pointer;
        padding: 0;
    }
</style>
