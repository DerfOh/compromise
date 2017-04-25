package edu.umflint.superteam.compromise.API;

import android.os.AsyncTask;
import android.util.Log;
import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.ProtocolException;
import java.net.URL;

/**
 * Created by nsammut on 3/26/17.
 */

    public class HttpClient extends AsyncTask<Void, Void, Boolean> {

    private final String mMethod;
    private final String mUrl;
    private final String mParameters;
    private StringBuffer response;

    public HttpClient(String method, String url, String parameters) {
        mMethod = method;
        mUrl = url;
        mParameters = parameters;
    }

    @Override
    protected Boolean doInBackground(Void... params) {
        try {

            URL obj = new URL(mUrl);
            HttpURLConnection con = (HttpURLConnection) obj.openConnection();
            con.setRequestMethod(mMethod);

            if(!mMethod.equals("GET")) {
                con.setRequestProperty("Content-Type", "application/x-www-form-urlencoded");
                String urlParameters = mParameters;
                con.setDoOutput(true);
                DataOutputStream wr = new DataOutputStream(con.getOutputStream());
                wr.writeBytes(urlParameters);
                wr.flush();
                wr.close();
            }

            int responseCode = con.getResponseCode();
            Log.i("HttpClient", "Sending '" + mMethod + "' request to URL : " + mUrl);
            Log.i("HttpClient", "Response Code : " + responseCode);

            BufferedReader in = new BufferedReader(new InputStreamReader(con.getInputStream()));
            String inputLine;
            response = new StringBuffer();

            while ((inputLine = in.readLine()) != null) {
                response.append(inputLine);
            }
            in.close();
        } catch (ProtocolException e) {
            e.printStackTrace();
        } catch (MalformedURLException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
        Log.i("HttpClient", response.toString());
        return true;
    }
}