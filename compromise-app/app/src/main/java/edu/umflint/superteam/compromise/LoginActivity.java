package edu.umflint.superteam.compromise;

import android.content.ActivityNotFoundException;
import android.content.Intent;
import android.content.SharedPreferences;
import android.preference.PreferenceManager;
import android.support.v7.app.AppCompatActivity;

import android.os.AsyncTask;

import android.os.Bundle;
import android.text.TextUtils;
import android.util.Log;
import android.view.KeyEvent;
import android.view.View;
import android.view.View.OnClickListener;
import android.view.inputmethod.EditorInfo;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.Properties;
import java.util.concurrent.ExecutionException;

import javax.mail.Authenticator;
import javax.mail.Message;
import javax.mail.MessagingException;
import javax.mail.PasswordAuthentication;
import javax.mail.Session;
import javax.mail.Transport;
import javax.mail.internet.AddressException;
import javax.mail.internet.InternetAddress;
import javax.mail.internet.MimeMessage;

import edu.umflint.superteam.compromise.API.GetPassword;
import edu.umflint.superteam.compromise.API.GetTasks;

/**
 * A login screen that offers login via email/password.
 */
public class LoginActivity extends AppCompatActivity {

    Session session = null;
    /**
     * Keep track of the login task to ensure we can cancel it if requested.
     */
    private UserLoginTask mAuthTask = null;
    // UI references.
    private EditText mEmailView;
    private EditText mPasswordView;
    private String email;
    private String password;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_login);

        SharedPreferences prefs = PreferenceManager.getDefaultSharedPreferences(getApplicationContext());
        if(!prefs.getString("EmailAddress", null).isEmpty())
        {
            Intent intent = new Intent(LoginActivity.this, MainActivity.class);
            startActivity(intent);
        }

        // Set up the login form.
        mEmailView = (EditText) findViewById(R.id.email);
        mPasswordView = (EditText) findViewById(R.id.password);
        Button contactUs = (Button) findViewById(R.id.contact_us_btn);
        contactUs.setOnClickListener(new OnClickListener(){
            @Override
            public void onClick(View view) {
                try {
                    Intent emailIntent = new Intent(Intent.ACTION_SEND);
                    emailIntent.putExtra(Intent.EXTRA_EMAIL, new String[] { "Compromise@nicholassammut.com" });
                    emailIntent.putExtra(Intent.EXTRA_SUBJECT, "Compromise Android Application Feedback");
                    emailIntent.setType("message/rfc822");
                    startActivity(emailIntent);
                } catch(ActivityNotFoundException anfe) {
                    Toast.makeText(LoginActivity.this, "No Email Client Installed!", Toast.LENGTH_SHORT).show();
                }
            }
        });

        Button lostPassword = (Button) findViewById(R.id.lost_password);
        lostPassword.setOnClickListener(new OnClickListener(){
            @Override
            public void onClick(View view) {
                Properties props = new Properties();
                props.put("mail.smtp.host", "mail.nicholassammut.com");
                props.put("mail.smtp.socketFactory.port","465");
                props.put("mail.smtp.socketFactory.class","javax.net.ssl.SSLSocketFactory");
                props.put("mail.smtp.auth","true");
                props.put("mail.smtp.port","465");
                email = mEmailView.getText().toString();
                try {
                     password = new GetPassword(email).execute().get();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                } catch (ExecutionException e) {
                    e.printStackTrace();
                }

                session = Session.getDefaultInstance(props, new Authenticator() {
                    protected PasswordAuthentication getPasswordAuthentication() {
                        return new PasswordAuthentication("Compromise@nicholassammut.com", "Compromise123!");
                    }
                });
                RetrieveFeedTask task = new RetrieveFeedTask();
                task.execute(mEmailView.getText().toString());
            }

        class RetrieveFeedTask extends AsyncTask<String, Void, String> {
            @Override
            protected String doInBackground(String... params) {
                try {
                    Message message = new MimeMessage(session);
                    message.setFrom(new InternetAddress("Compromise@nicholassammut.com"));
                    message.setRecipients(Message.RecipientType.TO, InternetAddress.parse(params[0]));
                    message.setSubject("Compromise Android Application Password");
                    message.setContent("Your password is " + password + "!", "text/html ; charset=utf-8");
                    if(password.equals("\"\"")) {
                    }
                    else {
                        Transport.send(message);
                    }
                } catch (AddressException e) {
                    e.printStackTrace();
                } catch (MessagingException e) {
                    e.printStackTrace();
                }
                return password;
            }
            @Override
            protected void onPostExecute(String result) {
                if(password.equals("\"\"")) {
                    Toast.makeText(LoginActivity.this, "There was a problem! Check your Email Address!", Toast.LENGTH_SHORT).show();
                }
                else {
                    Toast.makeText(LoginActivity.this, "Your Password Has Been Sent!", Toast.LENGTH_SHORT).show();
                }


            }
        }
    });
        Button mEmailSignInButton = (Button) findViewById(R.id.signin_btn);
        mEmailSignInButton.setOnClickListener(new OnClickListener() {
            @Override
            public void onClick(View view) {
            attemptLogin();
            }
        });

    }

    /**
     * Attempts to sign in or register the account specified by the login form.
     * If there are form errors (invalid email, missing fields, etc.), the
     * errors are presented and no actual login attempt is made.
     */
    private boolean attemptLogin() {

        // Reset errors.
        mEmailView.setError(null);
        mPasswordView.setError(null);

        // Store values at the time of the login attempt.
        String email = mEmailView.getText().toString();
        String password = mPasswordView.getText().toString();

        boolean cancel = false;
        View focusView = null;

        // Check for a valid password, if the user entered one.
        if (TextUtils.isEmpty(password)) {
            mPasswordView.setError(getString(R.string.error_invalid_password));
            focusView = mPasswordView;
            cancel = true;
        }

        // Check for a valid email address.
        if (TextUtils.isEmpty(email)) {
            mEmailView.setError(getString(R.string.error_field_required));
            focusView = mEmailView;
            cancel = true;
        } else if (!isEmailValid(email)) {
            mEmailView.setError(getString(R.string.error_invalid_email));
            focusView = mEmailView;
            cancel = true;
        }

        if (cancel) {
            // There was an error; don't attempt login and focus the first
            // form field with an error.
            focusView.requestFocus();
        } else {
            // Show a progress spinner, and kick off a background task to
            // perform the user login attempt.
            try{
                mAuthTask = new UserLoginTask(email, password);
                mAuthTask.execute((Void) null);
                return true;
            }catch (Exception ex) {

            }
        }
        return true;
    }

    private boolean isEmailValid(String email) {
        //TODO: Replace this with your own logic
        return email.contains("@");
    }

    /**
     * Represents an asynchronous login/registration task used to authenticate
     * the user.
     */
    public class UserLoginTask extends AsyncTask<Void, Void, Boolean> {

        private final String mEmail;
        private final String mPassword;
        private StringBuffer response;

        UserLoginTask(String email, String password) {
            mEmail = email;
            mPassword = password;
        }

        @Override
        protected Boolean doInBackground(Void... params) {
            // TODO: attempt authentication against a network service.

            try {
                String url = "https://api.compromise.rocks/api/auth/";
                URL obj = new URL(url);
                HttpURLConnection con = (HttpURLConnection) obj.openConnection();

                //add request header
                con.setRequestMethod("POST");
                con.setRequestProperty("Content-Type", "application/x-www-form-urlencoded");

                String urlParameters = "EmailAddress=" + mEmail + "&Password=" + mPassword;

                // Send post request
                con.setDoOutput(true);
                DataOutputStream wr = new DataOutputStream(con.getOutputStream());
                wr.writeBytes(urlParameters);
                wr.flush();
                wr.close();

                int responseCode = con.getResponseCode();
                Log.i("Login", "Sending 'POST' request to URL : " + url);
                Log.i("Login", "Post parameters : " + urlParameters);
                Log.i("Login", "Response Code : " + responseCode);

                BufferedReader in = new BufferedReader(
                        new InputStreamReader(con.getInputStream()));
                String inputLine;
                response = new StringBuffer();

                while ((inputLine = in.readLine()) != null) {
                    response.append(inputLine);
                }
                in.close();
            } catch (Exception ex) {
                Log.e("Login", ex.toString());
                return false;
            }

                //print result
                Log.i("Login", response.toString());

                if(response.toString().equals("true"))
                {
                    return true;
                } else {
                    return false;
                }
        }

        @Override
        protected void onPostExecute(final Boolean success) {
            mAuthTask = null;

            if (success) {
                Toast.makeText(getApplication().getBaseContext(), "Login Successful!", Toast.LENGTH_SHORT).show();
                Intent intent = new Intent(LoginActivity.this, MainActivity.class);
                SharedPreferences prefs = PreferenceManager.getDefaultSharedPreferences(getApplicationContext());
                SharedPreferences.Editor editor = prefs.edit();
                editor.putString("EmailAddress", mEmail);
                editor.putInt("SelectedGroup", -1);
                editor.apply();
                startActivity(intent);
                finish();
            } else {
                Toast.makeText(getApplication().getBaseContext(), "Login Failed!", Toast.LENGTH_SHORT).show();
                mPasswordView.setError(getString(R.string.error_incorrect_password));
                mPasswordView.requestFocus();
            }
        }

        @Override
        protected void onCancelled() {
            mAuthTask = null;
        }
    }
}

