<script lang="ts">
  import { fade } from "svelte/transition";
  import Text from "./text/Text.svelte";
  import backgroundImg from "../../assets/venue/background.png";
  import locationImg from "../../assets/venue/location.svg";
  import calendarImg from "../../assets/venue/calendar.svg";
  import dressingImg from "../../assets/venue/dressing.svg";
  import "add-to-calendar-button";
  import {
    atcb_action,
    type ATCBActionEventConfig,
  } from "add-to-calendar-button";
  import { getTextDefault } from "./text/text";

  let {
    closeCb,
  }: {
    closeCb: () => void;
  } = $props();

  // location
  const locationURL = getTextDefault("location-link", "/#");
  const dressingURL = getTextDefault("dressing-link", "/#");

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

<div class="venue-wrapper" transition:fade={{ duration: 200 }}>
  <div class="closable">
    <button class="close" onclick={closeCb}>&times;</button>
    <div class="venue-container">
      <div class="parents formal">
        <p class="middle"><Text key="venue-parentsmiddle" /></p>
        <div class="parents-side">
          <p class="left"><Text key="venue-parentsleft" /></p>
          <p class="right"><Text key="venue-parentsright" /></p>
        </div>
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

        <a href={locationURL} target="_blank" class="action">
          <img src={locationImg} alt="maps" class="action-img" />
          <span class="action-description formal">
            <Text key="venue-locationdescription" />
          </span>
          <span class="action-call formal">
            <Text key="venue-locationcall" />
          </span>
        </a>

        <a href={dressingURL} target="_blank" class="action">
          <img src={dressingImg} alt="dressing" class="action-img" />
          <span class="action-description formal">
            <Text key="venue-dressingdescription" />
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

    border-radius: 1rem;
    border: 1px solid transparent;
  }

  .parents {
    width: 100%;
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;

    color: #dedacda0;
    font-style: italic;
  }

  .parents-side {
    display: flex;
    width: 90%;
    flex-flow: row nowrap;
    justify-content: space-between;
  }

  .left {
    text-align: left;
  }

  .middle {
    text-align: center;
  }

  .right {
    text-align: right;
  }

  .name {
    text-align: center;
    color: #dedacdff;
  }

  .description {
    font-style: italic;
    text-align: center;

    color: #dedacdc0;
  }

  .actions {
    width: 100%;
    display: flex;
    align-items: center;
  }

  .action {
    margin: 0;
    border: 0;
    padding: 0;
    background: transparent;
    text-decoration: none;
    flex: 1 1 0px;
    box-sizing: border-box;

    text-align: center;
    color: #dedacdff;

    display: flex;
    flex-flow: column nowrap;
    justify-content: start;
    align-items: center;

    border: 1px solid transparent;
    border-radius: 1rem;
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
    font-style: italic;
  }

  .footer {
    display: flex;
    flex-flow: column;
    justify-content: start;
    align-items: center;
  }

  .footer-line {
    border-top: 0.1rem solid #dedacdff;
  }

  .footer-text {
    text-align: center;
    font-style: italic;
    color: #dedacda0;
  }

  p {
    margin: 0;
  }

  @media only screen and (max-width: 768px) {
    .parents {
      height: 5rem;
      font-size: 0.9rem;
      gap: 0.5rem;
    }
    .parents-side {
      width: 100%;
    }

    .name {
      font-size: 6rem;
      line-height: 5rem;
      margin-top: 3rem;
    }
    .description {
      font-size: 1.2rem;
      margin-bottom: 2rem;
    }
    .actions {
      flex-flow: column nowrap;
      justify-content: center;
      align-items: center;
    }
    .action {
      width: 400px;
      padding: 1rem;
      flex-flow: row nowrap;
      justify-content: flex-start;
      gap: 2rem;
    }
    .action-call {
      font-size: 0.8rem;
    }
    .action-img {
      height: 2.5rem;
    }
    .action-description {
      font-size: 1rem;
    }
    .footer {
      margin-top: 1.5rem;
    }
    .footer-line {
      width: 3rem;
      height: 0.5rem;
    }
    .footer-text {
      font-size: 1rem;
    }
  }

  @media only screen and (min-width: 768px) {
    .parents {
      height: 4rem;
      font-size: 1.1rem;
    }
    .name {
      font-size: 6rem;
      line-height: 5rem;
      margin-top: 3rem;
    }
    .description {
      font-size: 1.2rem;
      margin-bottom: 4rem;
    }
    .actions {
      height: 8rem;
      flex-flow: row nowrap;
      justify-content: space-around;
    }
    .action {
      width: 30%;
      padding: 1rem;
    }
    .action-call {
      font-size: 0.8rem;
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
    }
    .footer-line {
      width: 3rem;
      height: 0.5rem;
    }
    .footer-text {
      font-size: 1rem;
    }
  }
</style>
