<script lang="ts">
  import "add-to-calendar-button";
  import {
    atcb_action,
    type ATCBActionEventConfig,
  } from "add-to-calendar-button";
  import { getTextDefault } from "../text/text";

  import iconCalendar from "../../assets/venue/icon-calendar.png";

  const name = getTextDefault("calendar-name", "Wedding");
  const label = getTextDefault("calendar-label", "Calendar");
  const location = getTextDefault("calendar-location", "Brazil");
  const date = getTextDefault("calendar-date", "2027-01-01");
  const startTime = getTextDefault("calendar-starttime", "15:00");
  const endTime = getTextDefault("calendar-endtime", "21:00");

  const config: ATCBActionEventConfig = {
    name: name,
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

  let button: HTMLElement | null = $state(null);

  const open = () => {
    if (button) {
      atcb_action(config, button);
    }
  };
</script>

<button id="btn-calendar" onclick={open} bind:this={button}>
  <img src={iconCalendar} alt={label} />
</button>

<!-- <add-to-calendar-button
  {label}
  {name}
  {location}
  startDate={date}
  {startTime}
  {endTime}
  options="'Apple','Google','iCal'"
  listStyle="dropup-static"
  timeZone="America/Sao_Paulo"
  hideBackground
  hideIconButton
  hideCheckmark
  buttonStyle="round"
></add-to-calendar-button> -->

<style>
  #btn-calendar {
    background: transparent;
    border: none;
    cursor: pointer;
  }

  img {
    max-width: 50px;
  }
</style>
