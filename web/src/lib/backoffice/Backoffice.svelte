<script lang="ts">
  import Password from "./Password.svelte";
  import AttendeeTable from "./attendee/Table.svelte";
  import InviteTable from "./invite/Table.svelte";
  import type { Payload } from "./common";

  type ShowType = "attendee" | "invites" | "password";

  let payload: Payload | null = $state<Payload | null>(null);
  let show: ShowType = $state<ShowType>("password");

  const setContent = (x: Payload) => {
    payload = x;
    show = "attendee";
  };

  const setShow = (s: ShowType) => {
    show = s;
  };
</script>

<svelte:head>
  <link
    rel="stylesheet"
    href="https://assets.ubuntu.com/v1/vanilla_framework_version_4.59.0.min.css"
  />
</svelte:head>

<div class="l-main">
  <section class="p-strip is-shallow">
    <div class="row">
      <div class="col-12">
        {#if show !== "password"}
          <div class="p-tabs">
            <div class="p-tabs__list" role="tablist" aria-label="Attendees">
              <div class="p-tabs__item">
                <button
                  class="p-tabs__link"
                  role="tab"
                  onclick={(e) => setShow("attendee")}
                >
                  Attendees
                </button>
              </div>
              <div class="p-tabs__item">
                <button
                  class="p-tabs__link"
                  role="tab"
                  onclick={(e) => setShow("invites")}
                >
                  Invites
                </button>
              </div>
            </div>

            <div tabindex="0" role="tabpanel" hidden={show !== "attendee"}>
              <AttendeeTable content={payload?.attendees || []} />
            </div>

            <div tabindex="0" role="tabpanel" hidden={show !== "invites"}>
              <InviteTable content={payload?.invites || []} />
            </div>
          </div>
        {:else}
          <Password {setContent} />
        {/if}
      </div>
    </div>
  </section>
</div>
