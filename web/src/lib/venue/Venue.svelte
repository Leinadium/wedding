<script lang="ts">
  import { fade } from "svelte/transition";
  import Text from "../text/Text.svelte";
  import backgroundImg from "../../assets/venue/background.png";
  import locationImg from "../../assets/venue/location.svg";
  import calendarImg from "../../assets/venue/calendar.svg";
  import dressingImg from "../../assets/venue/dressing.svg";
  import "add-to-calendar-button";
  import {
    atcb_action,
    type ATCBActionEventConfig,
  } from "add-to-calendar-button";
  import { getText, getTextDefault } from "../text/text";

  let {
    closeCb,
  }: {
    closeCb: () => void;
  } = $props();

  // location
  const locationURL = getTextDefault("location-link", "/#");

  // calendar
  const nameCalendar = getTextDefault("calendar-name", "Wedding");
  const labelCalendar = getTextDefault("calendar-label", "Calendar");
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

  let buttonCalendar: HTMLElement | null = $state(null);

  const openCalendar = () => {
    if (buttonCalendar) {
      atcb_action(config, buttonCalendar);
    }
  };
</script>

<div class="venue-wrapper" transition:fade={{ duration: 300 }}>
  <div class="closable">
    <button class="close" onclick={closeCb}>&times;</button>
    <div class="venue-container" in:fade={{ delay: 300, duration: 500 }}>
      <div class="parents formal">
        <p class="left"><Text key="venue-parentsleft" /></p>
        <p class="middle"><Text key="venue-parentsmiddle" /></p>
        <p class="right"><Text key="venue-parentsright" /></p>
      </div>

      <span class="name cursive">
        <Text key="venue-name" />
      </span>

      <span class="description formal">
        <Text key="venue-description" />
      </span>

      <div class="actions">
        <button
          type="button"
          class="action"
          onclick={openCalendar}
          bind:this={buttonCalendar}
        >
          <img src={calendarImg} class="action-img" alt={labelCalendar} />
          <span class="action-description formal">
            <Text key="venue-calendardescription" />
          </span>
          <span class="action-call formal">
            <Text key="venue-calendarcall" />
          </span>
        </button>
        <div class="action">
          <img src={dressingImg} alt="dressing" class="action-img" />
          <span class="action-description formal">
            <Text key="venue-dressingdescription" />
          </span>
        </div>
        <a href={locationURL} target="_blank" class="action">
          <img src={locationImg} alt="maps" class="action-img" />
          <span class="action-description formal">
            <Text key="venue-locationdescription" />
          </span>
          <span class="action-call formal">
            <Text key="venue-locationcall" />
          </span>
        </a>
      </div>

      <div class="footer">
        <div class="footer-line"></div>
        <p class="footer-text formal">
          <Text key="venue-footer" />
        </p>
      </div>
    </div>
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

  .closable {
    width: calc(80vw + 2rem);
    max-width: calc(800px + 2rem);
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
    padding: 1rem;
  }

  .venue-container {
    box-sizing: border-box;
    position: relative;
    width: 100%;
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
    justify-content: start;
    align-items: center;
  }

  .close {
    align-self: flex-end;
    background: none;
    border: none;
    font-size: 2rem;
    color: #dedacdff;
    cursor: pointer;
    z-index: 10;
  }

  .parents {
    width: 100%;
    height: 4rem;
    display: flex;
    flex-flow: row nowrap;
    justify-content: space-between;

    color: #dedacda0;

    font-style: italic;
    font-size: 1.1rem;
  }

  .left {
    align-self: flex-end;
    text-align: left;
  }

  .middle {
    align-self: flex-start;
    text-align: center;
  }

  .right {
    align-self: flex-end;
    text-align: right;
  }

  .name {
    font-size: 6rem;
    line-height: 5rem;
    text-align: center;
    margin-top: 3rem;

    color: #dedacdff;
  }

  .description {
    font-size: 1.2rem;
    font-style: italic;
    text-align: center;
    margin-bottom: 4rem;

    color: #dedacdc0;
  }

  .actions {
    width: 100%;
    height: 8rem;
    display: flex;
    flex-flow: row nowrap;
    justify-content: space-around;
    align-items: center;
  }

  .action {
    margin: 0;
    border: 0;
    padding: 0;
    background: transparent;
    text-decoration: none;

    width: 30%;
    flex: 1 1 0px;

    text-align: center;
    color: #dedacdff;

    display: flex;
    flex-flow: column nowrap;
    justify-content: start;
    align-items: center;

    border: 1px solid transparent;
    border-radius: 1rem;
    padding: 1rem;
  }

  button,
  a {
    cursor: pointer;
  }

  button:hover,
  a:hover {
    border: 1px solid #dedacdff;
  }

  .action-call {
    font-size: 0.8rem;
    font-style: italic;
  }

  .action-img {
    height: 2rem;
    margin-bottom: 0.3rem;
  }

  .action-description {
    font-size: 1.2rem;
  }

  .footer {
    margin-top: 3rem;
    display: flex;
    flex-flow: column;
    justify-content: start;
    align-items: center;
  }

  .footer-line {
    width: 3rem;
    height: 0.5rem;
    border-top: 0.1rem solid #dedacdff;
  }

  .footer-text {
    font-size: 1rem;
    text-align: center;
    font-style: italic;
    color: #dedacda0;
  }

  p {
    margin: 0;
  }
</style>
