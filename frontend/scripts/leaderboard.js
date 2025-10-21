function getRankClass(index) {
  switch (index) {
    case 0:
      return "first";
    case 1:
      return "second";
    case 2:
      return "third";
    default:
      return "";
  }
}

function getRankDisplay(index) {
  switch (index) {
    case 0:
      return "🥇";
    case 1:
      return "🥈";
    case 2:
      return "🥉";
    default:
      return `#${index + 1}`;
  }
}

function createTable(data, type) {
  if (!data || data.length === 0) {
    return `<div class="error">No ${type} found in leaderboard.</div>`;
  }

  let tableHTML = `
        <table>
            <thead>
                <tr>
                    <th>Rank</th>
                    ${type === "users" || type === "users-units" ? "<th>User</th>" : "<th>Item</th>"}
                    <th>${type === "users-units" ? "Units" : "Drank"}</th>
                </tr>
            </thead>
            <tbody>
    `;

  data.forEach((object, index) => {
    const rankClass = getRankClass(index);
    const rankDisplay = getRankDisplay(index);

    tableHTML += `
            <tr>
                <td class="rank ${rankClass}">${rankDisplay}</td>
                <td>${type === "users" || type === "users-units" ? (object.username ? `<a href="./user.html?username=${object.username}">${object.username}</a>` : "Unknown User") : object.name ? `<a href="./item.html?name=${encodeURIComponent(object.name)}">${object.name}</a>` : "Unknown Item"}</td>
                <td>${type === "users" ? object.consumed : type === "items" ? object.consumed : object.units}</td>
            </tr>
        `;
  });

  tableHTML += `
            </tbody>
        </table>
    `;

  return tableHTML;
}

function formatDateReadable(date) {
  const weekday = new Intl.DateTimeFormat("en-GB", { weekday: "long" }).format(
    date,
  );
  const month = new Intl.DateTimeFormat("en-GB", { month: "long" }).format(
    date,
  );
  const day = date.getDate();

  const getSuffix = (d) => {
    if (d > 3 && d < 21) return "th";
    switch (d % 10) {
      case 1:
        return "st";
      case 2:
        return "nd";
      case 3:
        return "rd";
      default:
        return "th";
    }
  };

  const dayWithSuffix = day + getSuffix(day);

  return `${weekday} ${dayWithSuffix} ${month}`;
}
