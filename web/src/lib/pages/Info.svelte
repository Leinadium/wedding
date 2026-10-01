<script lang="ts">
  import headerImg from "../../assets/info/header.jpg";
  import churchImg from "../../assets/info/church.svg";
  import carImg from "../../assets/info/car.svg";
  import dressingImg from "../../assets/info/dressing.svg";
  import dressImg from "../../assets/info/dress.svg";
  import ringImg from "../../assets/info/ring.svg";
  import finalImg from "../../assets/info/final.png";

  import { Page } from "../common";
  import Header from "../pages/Header.svelte";
  import Text from "../text/Text.svelte";
  import { onDestroy, onMount } from "svelte";
  import BackButton from "./BackButton.svelte";
  import "add-to-calendar-button";
  import {
    atcb_action,
    type ATCBActionEventConfig,
  } from "add-to-calendar-button";
  import { getTextDefault } from "../text/text";

  let {
    setState,
  }: {
    setState: (page: Page) => void;
  } = $props();

  // location
  const locationURL = getTextDefault("location-link", "/#");
  const dressingURL = getTextDefault("dressing-link", "/#");

  // calendar
  const nameCalendar = getTextDefault("calendar-name", "Wedding");
  const location = getTextDefault("calendar-location", "Brazil");
  const date = getTextDefault("calendar-date", "2027-01-01");
  const startTime = getTextDefault("calendar-starttime", "15:00");
  const endTime = getTextDefault("calendar-endtime", "21:00");

  const config: ATCBActionEventConfig = {
    name: nameCalendar,
    location: location,
    startDate: date,
    startTime: startTime,
    endTime: endTime,
    options: ["Apple", "Google", "iCal"],
    listStyle: "modal",
    timeZone: "America/Sao_Paulo",
    hideBackground: true,
    hideCheckmark: true,
    buttonStyle: "round",
  };

  const openCalendar = (e: Event) => {
    e.preventDefault();
    var x = document.getElementById("calendar");
    if (x) {
      atcb_action(config, x);
    }
  };

  function back() {
    setState(Page.Landing);
  }

  function scrollTop() {
    window.scrollTo({ top: 0, behavior: "smooth" });
  }

  onMount(scrollTop);
  onDestroy(scrollTop);
</script>

<div class="info">
  <Header src={headerImg} text={getTextDefault("info-title", "The Big Day")} />

  <div class="content">
    <span class="date-place cursive mini-title">
      <Text key="info-date-place" />
    </span>

    <span class="date formal">
      <Text key="info-date" />
    </span>

    <span class="place-name formal">
      <Text key="info-place-name" />
    </span>

    <span class="place-description formal">
      <Text key="info-place-description" />
    </span>

    <a
      id="calendar"
      href="/"
      onclick={openCalendar}
      class="place-button formal"
    >
      <Text key="info-place-button" />
    </a>

    <img class="church icon" src={churchImg} alt="church" />

    <span class="directions cursive mini-title">
      <Text key="info-directions" />
    </span>
    <img class="car icon" src={carImg} alt="car" />

    <span class="directions-title formal">
      <Text key="info-directions-map-title" />
    </span>
    <a href={locationURL} class="directions-description formal">
      <Text key="info-directions-map-description" />
    </a>

    <span class="directions-title formal">
      <Text key="info-directions-car-title" />
    </span>
    <span class="directions-description formal">
      <Text key="info-directions-car-description" />
    </span>

    <span class="directions-title formal">
      <Text key="info-directions-uber-title" />
    </span>
    <span class="directions-description formal">
      <Text key="info-directions-uber-description" />
    </span>

    <span class="dresscode cursive mini-title">
      <Text key="info-dresscode" />
    </span>
    <img class="dressing icon" src={dressingImg} alt="car" />

    <span class="dresscode-title formal">
      <Text key="info-dresscode-title" />
    </span>
    <a href={dressingURL} class="dresscode-description formal">
      <Text key="info-dresscode-description" />
    </a>

    <span class="dressing-obs formal">
      <Text key="info-dressing-obs-1" />
    </span>
    <img class="dress icon" src={dressImg} alt="dress" />

    <span class="dressing-obs formal">
      <Text key="info-dressing-obs-2" />
    </span>

    <span class="final-title cursive mini-title">
      <Text key="info-final-title" />
    </span>
    <img class="ring icon" src={ringImg} alt="ring" />

    <span class="quote formal">
      <Text key="info-quote" />
    </span>

    <img class="final" src={finalImg} alt="final" />
    <span class="final-small formal">
      <Text key="info-final-small" />
    </span>
    <span class="final-name cursive">
      <Text key="info-final-name" />
    </span>

    <BackButton textKey="story-back" action={back} color={"#f9f6f1"} />
  </div>
</div>

<style>
  .info {
    display: flex;
    flex-flow: column nowrap;
    align-items: center;

    min-height: 90vh;
    width: 100vw;

    color: #f9f6f1;
    font-size: 1.2rem;
    text-align: justify;
  }

  .mini-title {
    padding-top: 3rem;
    font-size: 4rem;
    line-height: 3rem;
  }

  .content {
    display: flex;
    flex-flow: column;
    justify-content: flex-start;
    align-items: center;

    text-align: center;
  }

  a {
    text-decoration: none;
    color: #f9f6f1 !important;
  }

  .icon {
    width: 7rem;
  }

  .date {
    padding-top: 2rem;
  }

  .place-name {
    margin-top: 2rem;
  }

  .place-button {
    margin-top: 2rem;
  }

  .church {
    margin-top: 2rem;
  }

  .car {
    margin-top: 1.3rem;
  }

  .directions-title {
    font-weight: bold;
    font-size: 1.3rem;
    margin-top: 1.5rem;
  }
  .dresscode-title {
    margin-top: 1.3rem;
    font-size: 1.3rem;
    font-weight: bold;
  }

  .dressing-obs {
    margin-top: 1rem;
  }

  .final-title {
    margin-top: 1rem;
    font-size: 3.5rem;
  }
  .quote {
    margin-top: 1rem;
    font-size: 1.4rem;
  }

  .final {
    margin-top: 3rem;
    width: 70vw;
    max-width: 28rem;
  }

  .final-small {
    margin-top: 1rem;
  }

  .final-name {
    font-size: 3rem;
    margin-bottom: 3rem;
  }
</style>
