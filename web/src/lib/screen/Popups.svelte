<script lang="ts">
  let rsvp = $state(false);
  let invite = $state(false);

  function toggleRSVP() {
    rsvp = !rsvp;
  }

  function toggleInvite() {
    invite = !invite;
  }
</script>

<!-- RSVP wrapper -->
<div class="collapsible-wrapper">
  <button onclick={toggleRSVP} class="toggle-btn" aria-expanded={rsvpOpen}>
    RSVP {rsvpOpen ? "-" : "+"}
  </button>

  {#if rsvpOpen}
    <Invite closeCb={toggleRSVP} />
  {/if}
</div>
