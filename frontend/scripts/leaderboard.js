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
                    ${type === "users" || type === "users-by-units" ? "<th>User</th>" : "<th>Item</th>"}
                    <th>${type === "users-by-units" ? "Units" : "Drank"}</th>
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
                <td>${type === "users" || type === "users-by-units" ? (object.username ? `<a href="./user.html?username=${object.username}">${object.username}</a>` : "Unknown User") : object.name ? `<a href="./item.html?name=${object.name}">${object.name}</a>` : "Unknown Item"}</td>
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

