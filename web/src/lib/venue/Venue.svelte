<script lang="ts">
  import { fade } from "svelte/transition";
  import background from "../../assets/venue/background.png";
  import logo from "../../assets/venue/logo.png";
  import location from "../../assets/venue/location.svg";
  import calendar from "../../assets/venue/calendar.svg";
  import dressing from "../../assets/venue/dressing.svg";

  let {
    closeCb,
  }: {
    closeCb: () => void;
  } = $props();
</script>

<div class="venue-wrapper" transition:fade={{ duration: 300 }}>
  <div class="venue-container">
    <button class="close" onclick={closeCb}>&times;</button>

    <img id="logo" src={logo} alt="logo" />

    <div class="text-container">
      <img class="icon" src={calendar} alt="calendar" />
      <p class="formal">
        17 de Janeiro de 2027<br />
        Início da Cerimônia: 15h00
      </p>
    </div>

    <div class="text-container">
      <img class="icon" src={location} alt="location" />
      <p class="formal">
        Chateau do Lago<br />
        Estrada do Açude, 707<br />
        Alto da Boa Vista, Rio de Janeiro<br />
        20531-330
      </p>
    </div>

    <div class="text-container">
      <img class="icon" src={dressing} alt="dressing" />
      <p class="formal">Traje: Passeio Completo</p>
    </div>
    <div id="footer"></div>
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
    width: 50%;
    max-width: 800px;
    max-height: 90vh;
    overflow-y: auto;

    border-radius: 8px;
    padding: 1rem 1rem 1rem 1rem;
    box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);

    background-image: url("../../assets/venue/background.png");
    background-position: center;
    background-size: cover;

    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
    gap: 3rem;
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

  #logo {
    max-width: 3rem;
    height: auto;
  }

  .text-container {
    width: 100%;

    display: flex;
    flex-flow: column;
    justify-content: start;
    align-items: center;
    gap: 0.5rem;

    color: #f0f0f0;
  }

  .icon {
    max-width: 2rem;
    height: auto;
  }

  p {
    margin: 0;
    font-weight: 500;
    text-align: center;
  }
</style>
