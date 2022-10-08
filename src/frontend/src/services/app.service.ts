import DigQueryModel from "../models/digQuery";
import DigResponse from "../models/digResponse";

export class AppService {
    public async Dig(query: DigQueryModel): Promise<DigResponse> {
        const response = await fetch(`${process.env.REACT_APP_API_URL}api/v1/dig/`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(query)
        })
        return await response.json();
    }
}