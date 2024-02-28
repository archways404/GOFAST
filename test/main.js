const cheerio = require('cheerio');

async function getSchedule(url) {
	try {
		const response = await fetch(url);
		if (!response.ok) {
			throw new Error('Network response was not ok');
		}
		const html = await response.text();
		const $ = cheerio.load(html);
		const data = [];

		$('.data-grey, .data-white').each((index, row) => {
			const cells = $(row).find('td');
			const dateText = cells.eq(2).text().trim();
			const time = cells.eq(3).text().trim();
			let date = '';
			const dateMatch = dateText.match(/\d+/);
			if (dateMatch) {
				date = dateMatch[0].padStart(2, '0');
			}
			if (date.length === 1) {
				date = `0${date}`;
			}
			// Check if time is in the expected format
			const timeMatch = time.match(/(\d{2}:\d{2})-(\d{2}:\d{2})/);
			if (timeMatch) {
				const startTime = timeMatch[1];
				const endTime = timeMatch[2];
				const [startHour, startMinutes] = startTime.split(':').map(Number);
				const [endHour, endMinutes] = endTime.split(':').map(Number);
				const startInMinutes = startHour * 60 + startMinutes;
				const endInMinutes = endHour * 60 + endMinutes;
				const timeTotalInMinutes = endInMinutes - startInMinutes;
				const timeTotal = timeTotalInMinutes / 60;
				if (!isNaN(timeTotal)) {
					const formattedData = `${date},${timeTotal}`;
					data.push(formattedData);
				}
			}
		});
		return data;
	} catch (error) {
		console.error(error);
	}
}

let url =
	'https://schema.mau.se/setup/jsp/Schema.jsp?slutDatum=2024-03-31&sprak=SV&sokMedAND=true&startDatum=2024-02-29&moment=philip&resurser=k.BIT%20-%20IT';

async function main() {
	const schedule = await getSchedule(url);
	console.log(schedule);
}

console.time('getScheduleExecutionTime');
main().then(() => {
	console.timeEnd('getScheduleExecutionTime');
});
