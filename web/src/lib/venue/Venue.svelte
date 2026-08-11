<script lang="ts">
  import { fade } from "svelte/transition";
  import "add-to-calendar-button";
  import Calendar from "./Calendar.svelte";

  let {
    closeCb,
  }: {
    closeCb: () => void;
  } = $props();
</script>

<div class="venue-wrapper" transition:fade={{ duration: 300 }}>
  <div class="venue-container">
    <button class="close" onclick={closeCb}>&times;</button>

    <Calendar />
  </div>
</div>

<style>
  .venue-wrapper {
    position: fixed;
    top: 0;
    left: 0;

    width: 100vw;
    height: 100vh;

    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;

    background-color: rgba(0, 0, 0, 0.6);
    z-index: 999;
  }

  .venue-container {
    position: relative;
    width: 90%;
    max-width: 800px;
    max-height: 90vh;
    overflow-y: auto;

    border-radius: 8px;
    padding: 3rem 1rem 1rem 1rem;
    box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);

    background-image: url("../../assets/invite/texture.png");
    background-color: #f0f0f0;
    box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);

    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
  }

  .close {
    position: absolute;
    top: 1rem;
    right: 1rem;
    background: none;
    border: none;
    font-size: 1.5rem;
    color: #777;
    cursor: pointer;
    z-index: 10;
  }

  .cursive {
    font-family: "Italianno";
  }
</style>
