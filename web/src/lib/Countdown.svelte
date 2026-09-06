<script lang="ts">
  import { onDestroy, onMount } from "svelte";

  import backgroundImg from "../assets/venue/background.png";

  const weddingDate = new Date("2027-01-17T15:30:00"); // sample date

  let currentTime: Date = $state(new Date());
  let loop: number = $state(0);

  function updateTime() {
    currentTime = new Date();
  }

  let days: number = $state(0);
  let hours: number = $state(0);
  let minutes: number = $state(0);
  let seconds: number = $state(0);

  $effect(() => {
    const timeLeft = weddingDate.getTime() - currentTime.getTime();
    if (timeLeft < 0) return;
    days = Math.floor(timeLeft / (1000 * 60 * 60 * 24));
    hours = Math.floor((timeLeft % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
    minutes = Math.floor((timeLeft % (1000 * 60 * 60)) / (1000 * 60));
    seconds = Math.floor((timeLeft % (1000 * 60)) / 1000);
  });

  function formatTime(value: number): string {
    return value.toString().padStart(2, "0");
  }

  onMount(() => {
    currentTime = new Date();
    loop = setInterval(updateTime, 1000);
  });
  onDestroy(() => {
    clearInterval(loop);
  });
</script>

<div class="countdown">
  <div class="slot">
    <span class="number formal-num">{formatTime(days)}</span>
    <span class="text formal">days</span>
  </div>
  <!-- <span class="sep">:</span> -->
  <div class="slot">
    <span class="number formal-num">{formatTime(hours)}</span>
    <span class="text formal">hours</span>
  </div>
  <!-- <span class="sep">:</span> -->
  <div class="slot">
    <span class="number formal-num">{formatTime(minutes)}</span>
    <span class="text formal">minutes</span>
  </div>
  <!-- <span class="sep">:</span> -->
  <div class="slot">
    <span class="number formal-num">{formatTime(seconds)}</span>
    <span class="text formal">seconds</span>
  </div>
</div>

<style>
  .countdown {
    box-sizing: border-box;
    width: 100%;

    display: flex;
    flex-flow: row nowrap;

    color: #fae7b6;
    font-weight: bold;
    align-items: center;
  }

  .slot {
    flex-grow: 1;
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
  }

  .number {
    font-size: 2rem;
  }

  .text {
    font-size: 1rem;
  }
</style>
