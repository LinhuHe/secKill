package com.example.unittest;

import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;

public class Activity_two extends AppCompatActivity {

    private EditText ed_texttwo;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_two);
        ed_texttwo = findViewById(R.id.ed_Texttwo);
    }

    public void Onbtntwo(View view)
    {
        Toast.makeText(getApplication(),ed_texttwo.getText().toString(),Toast.LENGTH_LONG).show();
    }
}
